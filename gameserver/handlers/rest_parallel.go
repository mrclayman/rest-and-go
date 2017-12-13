package handlers

import (
	"net/http"

	"github.com/mrclayman/rest-and-go/gameserver/core"
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
	pool  RESTWorkerPool
	queue RESTJobQueue
	disp  *MainDispatcher
	quit  chan bool
}

// NewRESTWorker instantiates a new worker
// for REST interface jobs
func NewRESTWorker(p RESTWorkerPool, c *core.Core) RESTWorker {
	return RESTWorker{
		pool:  p,
		queue: make(RESTJobQueue),
		disp:  NewMainDispatcher(c),
		quit:  make(chan bool),
	}
}

// Start starts the processing loop of the worker
func (w *RESTWorker) Start() {
	go func() {
		for {
			// Register the worker as available for job processing
			w.pool <- w.queue

			// Wait for a job or a quit signal
			select {
			case job := <-w.queue:
				w.disp.ServeHTTP(job.Resp, job.Req)
			case <-w.quit:
				break
			}
		}
	}()
}

// RESTJobDispatcher picks jobs from the job
// queue and dispatches them to workers
type RESTJobDispatcher struct {
	pool           RESTWorkerPool
	queue          RESTJobQueue
	maxWorkerCount uint
}

// NewRESTJobDispatcher allocates a new REST job dispatcher
// object and returns a pointer to it
func NewRESTJobDispatcher(queue RESTJobQueue, maxWorkerCount uint) *RESTJobDispatcher {
	return &RESTJobDispatcher{
		pool:           make(RESTWorkerPool, maxWorkerCount),
		queue:          queue,
		maxWorkerCount: maxWorkerCount,
	}
}

// Start instantiates the required number of workers
// and starts their processing function
func (d *RESTJobDispatcher) Start(c *core.Core) {
	for i := uint(0); i < d.maxWorkerCount; i++ {
		w := NewRESTWorker(d.pool, c)
		w.Start()
	}

	go d.dispatch()
}

// dispatch dispatches a waiting job
// to a worker that is available
func (d *RESTJobDispatcher) dispatch() {
	for {
		select {
		case job := <-d.queue:
			go func(job RESTJob) {
				// Pick a worker from the worker queue
				worker := <-d.pool

				// Assign the job to it
				worker <- job
			}(job)
		}
	}
}
