package manager

import (
	"context"
	"sync"

	"github.com/elizarpif/logger"
	"github.com/elizarpif/worker-manager/service"
)

// WorkerManager for managing several workers
type WorkerManager struct {
	service         *service.Service
	workersCount    int
	newWorkersCount int

	wt *sync.WaitGroup // waitGroup for safety of the workers

	isActiveChan chan bool // chan for activate/deactivate worker manager (with buffer len=1)
	isWorkerDone chan bool // chan for manipulating workers (unbuffered)
	isClose      chan bool // chan for closing manager (unbuffered)
}

// NewWorkerManager creates new worker manager
func NewWorkerManager(srv *service.Service, workersCount int, isActive bool) *WorkerManager {
	// cannot be blocked until the channel is read, so buffered
	isActiveCh := make(chan bool, 1)
	if isActive {
		isActiveCh <- true
	}

	return &WorkerManager{
		service:         srv,
		wt:              &sync.WaitGroup{},
		isActiveChan:    isActiveCh,
		workersCount:    workersCount,
		newWorkersCount: workersCount,
		isClose:         make(chan bool),
	}
}

func (w *WorkerManager) SetNewWorkerCount(n int) {
	if n < 0 {
		return
	}

	w.newWorkersCount = n
}

// work starts workers
func (w *WorkerManager) work(ctx context.Context, isWorkerDone chan bool) {
	for i := 1; i <= w.workersCount; i++ {
		// do Add before start of goroutine to avoid the race condition
		w.wt.Add(1)

		// pass i into goroutine (i as an argument, i = 1,2..)
		go func(i int) {
			defer w.wt.Done()

			worker := &Worker{
				service: w.service,
				number:  i,
			}

			worker.Run(ctx, isWorkerDone)
		}(i)
	}
}

func (w *WorkerManager) done(ctx context.Context) {
	// say goroutine-workers to stop
	for i := 0; i < w.workersCount; i++ {
		w.isWorkerDone <- true
	}

	// wait while all goroutine-workers done their jobs
	w.wt.Wait()

	logger.GetLogger(ctx).Info("[worker-manager] worker pool deactivated")
}

// Process executes managing with workers
func (w *WorkerManager) Process(ctx context.Context) {
	log := logger.GetLogger(ctx)
	defer log.Infof("[worker-manager] out from select")

	for {
		select {
		case isActive := <-w.isActiveChan:
			log.Infof("[worker-manager] is_active = %v", isActive)

			if isActive {
				w.workersCount = w.newWorkersCount
				w.isWorkerDone = make(chan bool)
				w.work(ctx, w.isWorkerDone)
				break
			}

			// else, when isActive == false
			w.done(ctx)

		case <-w.isClose:
			w.done(ctx)

			// close all channels
			close(w.isClose)
			close(w.isWorkerDone)
			close(w.isActiveChan)

			return
		}
	}
}

// Close deactivates worker
func (w *WorkerManager) Close(ctx context.Context) error {
	// it is blocked until a read from this channel is performed
	// this changing will be read in Process()
	w.isClose <- true

	logger.GetLogger(ctx).Info("[worker-manager] worker pool was closed")
	return nil
}

// Deactivate deactivates workers
func (w *WorkerManager) Deactivate() {
	// this changing will be read in Process()
	w.isActiveChan <- false
}

// Activate activates workers
func (w *WorkerManager) Activate() {
	// this changing will be read in Process()
	w.isActiveChan <- true
}
