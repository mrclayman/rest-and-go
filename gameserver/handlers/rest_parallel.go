package handlers

import (
	"net/http"
)

// RESTJob defines a job coming
// from the server's REST interface
type RESTJob struct {
	Req  *http.Request
	Resp http.ResponseWriter
}

// RESTJobQueue defines a queue for
// REST jobs
type RESTJobQueue chan RESTJob

// RESTWorkerPool defines the type for
// a pool (channel) of REST job workers
type RESTWorkerPool chan chan RESTJob

// RESTWorker defines the type for a worker
// handling RESTJobs
type RESTWorker struct {
	Pool  RESTWorkerPool
	Queue RESTJobQueue
	quit  chan bool
}

// NewRESTWorker instantiates a new worker
// for REST interface jobs
func NewRESTWorker(pool RESTWorkerPool) RESTWorker {
	return RESTWorker{
		Pool:  pool,
		Queue: make(RESTJobQueue),
		quit:  make(chan bool),
	}
}

func (w RESTWorker) Start() {
	go func() {
		for {
			// Register the worker as available for job processing
			w.Pool <- w.Queue

			select {
			case job := <-w.Queue:
				// Time to do some work
			}
		}
	}()
}
