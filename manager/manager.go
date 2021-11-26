package manager

import (
	"context"
	"sync"

	"github.com/elizarpif/logger"
)

// WorkerManager for resorting by workers
type WorkerManager struct {
	workersCount int

	wt *sync.WaitGroup // waitGroup for safety of the workers

	isActiveChan chan bool // chan for activate/deactivate worker pool (with buffer len=1)
	isWorkerDone chan bool // chan for manipulating workers (unbuffered)
	isClose      chan bool // chan for closing manager (unbuffered)
}

// NewWorkerManager creates new worker pool for resorting
func NewWorkerManager(workersCount int, isActive bool) *WorkerManager {
	// нельзя блокироваться, пока не будет прочитан канал, поэтому буферизованный
	isActiveCh := make(chan bool, 1)
	if isActive {
		isActiveCh <- true
	}

	return &WorkerManager{
		wt:           &sync.WaitGroup{},
		isActiveChan: isActiveCh,
		workersCount: workersCount,
		isClose:      make(chan bool),
	}
}

// work starts workers
func (w *WorkerManager) work(ctx context.Context, isWorkerDone chan bool) {
	for i := 1; i <= w.workersCount; i++ {
		// осуществляем Add до начала горутины, чтобы избежать состояния гонки
		w.wt.Add(1)

		// передаем i в горутину (т. к. передается как аргумент, i = 1,2..)
		go func(i int) {
			defer w.wt.Done()

			worker := &Worker{
				number: i,
			}

			worker.Run(ctx, isWorkerDone)
		}(i)
	}
}

func (w *WorkerManager) done(ctx context.Context) {
	// говорим горутинам-воркерам остановиться
	for i := 0; i < w.workersCount; i++ {
		w.isWorkerDone <- true
	}

	// ожидаем пока все горутины-воркеры закончат работу
	w.wt.Wait()

	logger.GetLogger(ctx).Infof("worker manager deactivated")
}

// Process executes managing with workers
func (w *WorkerManager) Process(ctx context.Context) {
	for {
		select {
		case isActive := <-w.isActiveChan:
			logger.GetLogger(ctx).Infof("is_active = %v", isActive)

			if isActive {
				w.isWorkerDone = make(chan bool)
				w.work(ctx, w.isWorkerDone)
				continue
			}

			// else, when isActive == false
			w.done(ctx)

		case <-w.isClose:
			w.done(ctx)

			// закрываем все каналы
			close(w.isClose)
			close(w.isWorkerDone)
			close(w.isActiveChan)

			return
		}
	}
}

// Close deactivates worker
func (w *WorkerManager) Close() error {
	// блокируется, пока не будет осуществлено чтение из этого канала
	// this changing will be read in Process()
	w.isClose <- true

	logger.GetLogger(context.Background()).Infof("worker manager was closed")
	return nil
}

// Deactivate deactivates workers
func (w *WorkerManager) Deactivate() {
	// this changing will be read in Process()
	w.isActiveChan <- false
}

// Activate activates workers
func (w *WorkerManager) Activate(nWorkers int) {
	w.workersCount = nWorkers
	// this changing will be read in Process()
	w.isActiveChan <- true
}
