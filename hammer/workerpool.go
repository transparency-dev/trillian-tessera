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

import "context"

type worker interface {
	Run(ctx context.Context)
	Kill()
}

func newWorkerPool(factory func() worker) workerPool {
	workers := make([]worker, 0)
	pool := workerPool{
		workers: workers,
		factory: factory,
	}
	return pool
}

// workerPool contains a collection of _running_ workers.
type workerPool struct {
	workers []worker
	factory func() worker
}

func (p *workerPool) Grow(ctx context.Context) {
	w := p.factory()
	p.workers = append(p.workers, w)
	go w.Run(ctx)
}

func (p *workerPool) Shrink(ctx context.Context) {
	if len(p.workers) == 0 {
		return
	}
	w := p.workers[len(p.workers)-1]
	p.workers = p.workers[:len(p.workers)-1]
	w.Kill()
}
