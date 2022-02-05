package service

import (
	"context"
	"math/rand"
	"time"
)

type Service struct {
	n int
}

func (s *Service) Process(ctx context.Context) error {
	s.n++

	// imitation of process
	// it could be working with database
	// or processing messages from Kafka
	sleepTime := time.Duration(rand.Intn(5))
	time.Sleep(time.Second * sleepTime)
	// logger.GetLogger(ctx).Infof("[Process] %d", s.n)

	return nil
}
