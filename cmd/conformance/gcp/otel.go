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

	"go.opentelemetry.io/otel"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	mexporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric"
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"k8s.io/klog/v2"
)

// initOTel initialises the open telemetry support for metrics and tracing.
//
// Tracing is enabled with statistical sampling, with the probability passed in.
// Returns a shutdown function which should be called just before exiting the process.
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

	me, err := mexporter.New()
	if err != nil {
		klog.Exitf("Failed to create metric exporter: %v", err)
		return nil
	}
	// initialize a MeterProvider that periodically exports to the GCP exporter.
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(me)),
	)
	shutdownFuncs = append(shutdownFuncs, mp.Shutdown)
	otel.SetMeterProvider(mp)

	te, err := texporter.New()
	if err != nil {
		klog.Exitf("Failed to create trace exporter: %v", err)
		return nil
	}
	// initialize a TracerProvier that periodically exports to the GCP exporter.
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(traceFraction)),
		sdktrace.WithBatcher(te),
	)
	shutdownFuncs = append(shutdownFuncs, mp.Shutdown)
	otel.SetTracerProvider(tp)

	return shutdown
}
