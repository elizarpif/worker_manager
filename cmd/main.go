package main

import (
	"context"
	"time"

	"github.com/elizarpif/worker-manager/manager"
	"github.com/elizarpif/worker-manager/service"
)

func main() {
	ctx := context.Background()
	someService := &service.Service{}

	workersManager := manager.NewWorkerManager(someService, 4, true)
	go workersManager.Process(ctx)
	defer workersManager.Close(ctx)

	time.Sleep(time.Second * 5)

	workersManager.Deactivate()
	workersManager.SetNewWorkerCount(1)
	workersManager.Activate()

	time.Sleep(time.Second * 10)
}
