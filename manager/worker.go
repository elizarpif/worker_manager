package manager

import (
	"context"

	"github.com/elizarpif/logger"
	"github.com/elizarpif/worker-manager/service"
)

// Worker for resorting
type Worker struct {
	service *service.Service

	number int
}

// Run executes resorting by worker
func (w *Worker) Run(ctx context.Context, isDone chan bool) {
	log := logger.GetLogger(ctx)
	log.Infof("worker %d started", w.number)

	for {
		select {
		case <-isDone:
			log.Infof("worker %d done", w.number)
			return
		default:
			log.Infof("worker %d do smth", w.number)

			err := w.service.Process(ctx)
			if err != nil {
				log.Error(err)
			}
		}
	}
}