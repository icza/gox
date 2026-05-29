package gox

import (
	"sync"
	"sync/atomic"
)

// WorkerPool is a simple worker pool implementation that processes jobs concurrently,
// and job processing can request abortion.
//
// [WorkerPool.Start] must be called before enqueuing jobs, and [WorkerPool.Wait] must be called
// after [WorkerPool.Start] to wait for all enqueued jobs to be processed.
// Jobs can only be enqueued between the [WorkerPool.Start] and [WorkerPool.Wait] calls, using the [WorkerPool.Enqueue] method.
//
// The worker pool is re-startable: after [WorkerPool.Wait] returns, [WorkerPool.Start] can be called again to start a new round of job processing.
//
// NOTE: this WorkerPool API is not final, may change in the future.
type WorkerPool[Job any] struct {
	// Number of workers to process jobs concurrently.
	// Defaults to 1 if not set.
	WorkersCount int

	// HandleJob is a function responsible to handle (process) a single job.
	// true return value may be used to request abortion.
	// The caller may choose to stop enqueuing new jobs when [WorkerPool.AbortRequested] returns true.
	HandleJob func(job Job) (requestAbort bool)

	workersWg *sync.WaitGroup
	jobsCh    chan Job
	resultsCh chan bool

	abortRequested atomic.Bool
}

// Start launches internal goroutines of the worker pool. Returns immediately.
func (wp *WorkerPool[Job]) Start() {
	// Reset state, initialize internal channels and wait groups:
	wp.abortRequested.Store(false)
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

// AbortRequested reports if a job handler has requested abortion.
func (wp *WorkerPool[Job]) AbortRequested() bool {
	return wp.abortRequested.Load()
}

// Enqueue adds a job to the processing queue.
func (wp *WorkerPool[Job]) Enqueue(job Job) {
	wp.jobsCh <- job
}

// Wait signals the end of jobs, and blocks until all enqueued jobs have been processed.
// Note: a worker pool is re-startable, but [WorkerPool.Wait] must be called and only once after a [WorkerPool.Start] call.
func (wp *WorkerPool[Job]) Wait() {
	// First close the jobs channel (no more jobs can be enqueued)...
	close(wp.jobsCh)
	// ... and wait for the workers to end.
	wp.workersWg.Wait()
	// Now we can close the results channel, which will end the result gathering goroutine.
	close(wp.resultsCh)
}
