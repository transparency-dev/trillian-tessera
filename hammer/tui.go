// Copyright 2024 The Tessera authors. All Rights Reserved.
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
	"flag"
	"fmt"
	"time"

	movingaverage "github.com/RobinUS2/golang-moving-average"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"k8s.io/klog/v2"
)

func hostUI(ctx context.Context, hammer *Hammer) {
	grid := tview.NewGrid()
	grid.SetRows(5, 0, 10).SetColumns(0).SetBorders(true)
	// Status box
	statusView := tview.NewTextView()
	grid.AddItem(statusView, 0, 0, 1, 1, 0, 0, false)
	// Log view box
	logView := tview.NewTextView()
	logView.ScrollToEnd()
	logView.SetMaxLines(10000)
	grid.AddItem(logView, 1, 0, 1, 1, 0, 0, false)
	if err := flag.Set("logtostderr", "false"); err != nil {
		klog.Exitf("Failed to set flag: %v", err)
	}
	if err := flag.Set("alsologtostderr", "false"); err != nil {
		klog.Exitf("Failed to set flag: %v", err)
	}
	klog.SetOutput(logView)

	helpView := tview.NewTextView()
	helpView.SetText("+/- to increase/decrease read load\n>/< to increase/decrease write load\nw/W to increase/decrease workers")
	grid.AddItem(helpView, 2, 0, 1, 1, 0, 0, false)

	app := tview.NewApplication()
	interval := 500 * time.Millisecond
	ticker := time.NewTicker(interval)
	go func() {
		lastSize := hammer.tracker.LatestConsistent.Size
		maSlots := int((30 * time.Second) / interval)
		growth := movingaverage.New(maSlots)
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s := hammer.tracker.LatestConsistent.Size
				growth.Add(float64(s - lastSize))
				lastSize = s
				qps := growth.Avg() * float64(time.Second/interval)
				text := fmt.Sprintf("Read: %s\nWrite: %s\nTreeSize: %d (Δ %.0fqps over %ds)\nTime-in-queue: %s\nObserved-time-to-integrate: %s",
					hammer.readThrottle.String(),
					hammer.writeThrottle.String(),
					s,
					qps,
					time.Duration(maSlots*int(interval))/time.Second,
					formatMovingAverage(hammer.queueTime),
					formatMovingAverage(hammer.integrationTime),
				)
				statusView.SetText(text)
				app.Draw()
			}
		}
	}()
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case '+':
			klog.Info("Increasing the read operations per second")
			hammer.readThrottle.Increase()
		case '-':
			klog.Info("Decreasing the read operations per second")
			hammer.readThrottle.Decrease()
		case '>':
			klog.Info("Increasing the write operations per second")
			hammer.writeThrottle.Increase()
		case '<':
			klog.Info("Decreasing the write operations per second")
			hammer.writeThrottle.Decrease()
		case 'w':
			klog.Info("Increasing the number of workers")
			hammer.randomReaders.Grow(ctx)
			hammer.fullReaders.Grow(ctx)
			hammer.writers.Grow(ctx)
		case 'W':
			klog.Info("Decreasing the number of workers")
			hammer.randomReaders.Shrink(ctx)
			hammer.fullReaders.Shrink(ctx)
			hammer.writers.Shrink(ctx)
		}
		return event
	})
	// logView.SetChangedFunc(func() {
	// 	app.Draw()
	// })
	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}

func formatMovingAverage(ma *movingaverage.MovingAverage) string {
	aMin, _ := ma.Min()
	aMax, _ := ma.Max()
	aAvg := ma.Avg()
	return fmt.Sprintf("%.0fms/%.0fms/%.0fms (min/avg/max)", aMin, aAvg, aMax)
}
