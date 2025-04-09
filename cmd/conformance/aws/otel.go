// Copyright 2025 The Tessera authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"errors"

	"go.opentelemetry.io/contrib/detectors/aws/ec2"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"k8s.io/klog/v2"
)

// initOTel initialises the open telemetry support for metrics and tracing.
//
// Tracing is enabled with statistical sampling, with the probability passed in.
// Returns a shutdown function which should be called just before exiting the process.
//
// AWS requires that the ADOT collector is running on the local machine, listening on port 4317.
// See https://aws-otel.github.io/docs/getting-started/collector
func initOTel(ctx context.Context, traceFraction float64) func(context.Context) {
	var shutdownFuncs []func(context.Context) error
	// shutdown combines shutdown functions from multiple OpenTelemetry
	// components into a single function.
	shutdown := func(ctx context.Context) {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		if err != nil {
			klog.Errorf("OTel shutdown: %v", err)
		}
	}

	// Code below is mostly taken from the OTEL AWS documentation: https://aws-otel.github.io/docs/getting-started/go-sdk/manual-instr

	// Create and start new OTLP metric exporter
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure(), otlpmetricgrpc.WithEndpoint("localhost:4317"))
	if err != nil {
		klog.Exitf("Failed to create new OTLP metric exporter: %v", err)
	}
	mp := metric.NewMeterProvider(metric.WithReader(metric.NewPeriodicReader(metricExporter)))
	shutdownFuncs = append(shutdownFuncs, mp.Shutdown)
	otel.SetMeterProvider(mp)

	// Create and start new OTLP trace exporter
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint("localhost:4317"))
	if err != nil {
		klog.Exitf("Failed to create new OTLP trace exporter: %v", err)
	}

	// Instantiate a new EC2 Resource detector
	ec2ResourceDetector := ec2.NewResourceDetector()
	resource, err := ec2ResourceDetector.Detect(context.Background())
	if err != nil {
		klog.Exitf("Failed to detect EC2 resource: %v", err)
	}
	idg := xray.NewIDGenerator()

	// attach traceIDRatioBasedSampler to tracer provider
	// Associate resource with TracerProvider
	tp := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter),
		trace.WithIDGenerator(idg),
		trace.WithResource(resource),
		trace.WithSampler(trace.TraceIDRatioBased(traceFraction)),
	)
	shutdownFuncs = append(shutdownFuncs, tp.Shutdown)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(xray.Propagator{})

	return shutdown
}
