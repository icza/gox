package gox

import (
	"sync"
	"sync/atomic"
)

// WorkerPool is a simple worker pool implementation that processes jobs concurrently,
// and job processing can request abortion.
//
// Start() must be called before enqueuing jobs, and Wait() must be called after Start() to wait for all enqueued jobs to be processed.
// Jobs can only be enqueued between the Start() and Wait() calls, using the Enqueue() method.
//
// NOTE: this WorkerPool API is not final, may change in the future.
type WorkerPool[Job any] struct {
	// Number of workers to process jobs concurrently.
	// Defaults to 1 if not set.
	WorkersCount int

	// HandleJob is a function responsible to handle (process) a single job.
	// true return value may be used to request abortion.
	// The caller may choose to stop enqueuing new jobs when AbortRequested() returns true.
	HandleJob func(job Job) (requestAbort bool)

	workersWg *sync.WaitGroup
	jobsCh    chan Job
	resultsCh chan bool

	abortRequested atomic.Bool
}

// Start launches internal goroutines of the worker pool. Returns immediately.
func (wp *WorkerPool[Job]) Start() {
	// Initialize internal channels and wait groups:
	wp.workersWg = &sync.WaitGroup{}
	wp.jobsCh = make(chan Job)
	wp.resultsCh = make(chan bool)

	// Launch a pool of workers for concurrent processing:
	for range ForceMin(wp.WorkersCount, 1) {
		wp.workersWg.Go(func() {
			for job := range wp.jobsCh {
				wp.resultsCh <- wp.HandleJob(job)
			}
		})
	}

	// Gather results
	handleResultWg := &sync.WaitGroup{}
	handleResultWg.Go(func() {
		for result := range wp.resultsCh {
			if result {
				wp.abortRequested.Store(true)
			}
		}
	})
}

// AbortRequested reports if a job handler has requested abortion. It returns true if any of the processed jobs requested abortion.
func (wp *WorkerPool[Job]) AbortRequested() bool {
	return wp.abortRequested.Load()
}

// Enqueue adds a job to the processing queue.
func (wp *WorkerPool[Job]) Enqueue(job Job) {
	wp.jobsCh <- job
}

// Wait signals the end of jobs, and blocks until all enqueued jobs have been processed.
// Can only be called once, and only after Start() has been called.
func (wp *WorkerPool[Job]) Wait() {
	close(wp.jobsCh)
	wp.workersWg.Wait()
	close(wp.resultsCh)
}
