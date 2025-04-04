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
	"log"

	mexporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric"
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"k8s.io/klog/v2"
)

var (
	tp *sdktrace.TracerProvider
	mp *sdkmetric.MeterProvider
)

func initOTel() func(context.Context) {
	// Set up OTel trace and metric exporters
	texp, err := texporter.New()
	if err != nil {
		log.Fatalf("unable to set up tracing: %v", err)
	}
	tp = sdktrace.NewTracerProvider(sdktrace.WithBatcher(texp))
	otel.SetTracerProvider(tp)

	mexp, err := mexporter.New()
	if err != nil {
		klog.Exitf("Failed to create exporter: %v", err)
	}
	mp = sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(mexp)),
	)

	return func(ctx context.Context) {
		if err := tp.Shutdown(ctx); err != nil {
			klog.Errorf("Failed to shut down trace provider: %v", err)
		}
		if err := mp.Shutdown(ctx); err != nil {
			klog.Errorf("Failed to shut down meter provider: %v", err)
		}
	}
}
