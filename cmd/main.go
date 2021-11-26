package main

import (
	"context"
	"time"

	"github.com/elizarpif/worker-manager/manager"
)

func main() {
	workersManager := manager.NewWorkerManager(10, true)
	go workersManager.Process(context.Background())
	defer workersManager.Close()

	time.Sleep(time.Second * 3)
}
