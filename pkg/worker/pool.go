// Copyright 2021 Nitric Pty Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package worker

import (
	"fmt"
	"sync"
	"time"
)

type WorkerPool interface {
	// WaitForMinimumWorkers - A blocking method
	WaitForMinimumWorkers(timeout int) error
	GetWorkerCount() int
	GetWorker() (Worker, error)
	AddWorker(Worker) error
	RemoveWorker(Worker) error
	Monitor() error
}

type ProcessPoolOptions struct {
	MinWorkers int
	MaxWorkers int
}

// ProcessPool - A worker pool that represent co-located processes
type ProcessPool struct {
	minWorkers int
	maxWorkers int
	workerLock sync.Mutex
	workers    []Worker
	poolErr    chan error
}

func (p *ProcessPool) GetWorkerCount() int {
	p.workerLock.Lock()
	defer p.workerLock.Unlock()
	return len(p.workers)
}

// GetMinWorkers - return the minimum number of workers for this pool
func (p *ProcessPool) GetMinWorkers() int {
	return p.minWorkers
}

// GetMaxWorkers - return the maximum number of workers for this pool
func (p *ProcessPool) GetMaxWorkers() int {
	return p.maxWorkers
}

// Monitor - Blocks the current thread to supervise this worker pool
func (p *ProcessPool) Monitor() error {
	// Returns a pool error
	// In future we can catch this and attempt to create new workers to recover
	err := <-p.poolErr

	return err
}

// WaitForMinimumWorkers - Waits for the configured minimum number of workers to be available in this pool
func (p *ProcessPool) WaitForMinimumWorkers(timeout int) error {
	maxWaitTime := time.Duration(timeout) * time.Second
	// Longer poll times, e.g. 200 milliseconds results in slow lambda cold starts (15s+)
	pollInterval := time.Duration(15) * time.Millisecond

	var waitedTime = time.Duration(0)
	for {
		if p.GetWorkerCount() >= p.minWorkers {
			break
		} else {
			if waitedTime < maxWaitTime {
				time.Sleep(pollInterval)
				waitedTime += pollInterval
			} else {
				return fmt.Errorf("available workers below required minimum of %d, %d available, timedout waiting for more workers", p.minWorkers, p.GetWorkerCount())
			}
		}
	}

	return nil
}

// GetWorker - Retrieves a worker from this pool
func (p *ProcessPool) GetWorker() (Worker, error) {
	p.workerLock.Lock()
	defer p.workerLock.Unlock()

	if len(p.workers) > 0 {
		return p.workers[0], nil
	} else {
		return nil, fmt.Errorf("no workers available in this pool")
	}
}

// RemoveWorker - Removes the given worker from this pool
func (p *ProcessPool) RemoveWorker(wrkr Worker) error {
	p.workerLock.Lock()
	defer p.workerLock.Unlock()

	for i, w := range p.workers {
		if wrkr == w {
			p.workers = append(p.workers[:i], p.workers[i+1:]...)
			if len(p.workers) < p.minWorkers {
				p.poolErr <- fmt.Errorf("insufficient workers in pool, need minimum of %d, %d available", p.minWorkers, len(p.workers))
			}

			return nil
		}
	}

	return fmt.Errorf("worker does not exist in this pool")
}

// AddWorker - Adds the given worker to this pool
func (p *ProcessPool) AddWorker(wrkr Worker) error {
	p.workerLock.Lock()
	defer p.workerLock.Unlock()

	workerCount := len(p.workers)

	// Ensure we haven't reached the maximum number of workers
	if workerCount > p.maxWorkers {
		return fmt.Errorf("max worker capacity reached! cannot add more workers")
	}

	p.workers = append(p.workers, wrkr)

	return nil
}

// NewProcessPool - Creates a new process pool
func NewProcessPool(opts *ProcessPoolOptions) WorkerPool {
	if opts.MaxWorkers < 1 {
		opts.MaxWorkers = 1
	}

	return &ProcessPool{
		minWorkers: opts.MinWorkers,
		maxWorkers: opts.MaxWorkers,
		workerLock: sync.Mutex{},
		workers:    make([]Worker, 0),
		poolErr:    make(chan error),
	}
}
