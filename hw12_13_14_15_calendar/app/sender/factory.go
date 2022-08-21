package sender

import (
	"context"
	rmqconsumer "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/rmq/consumer"
	"time"

	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/app/sender/rmqworker"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/logger"
)

func CreateWorker(ctx context.Context, config Config, logger logger.Logger) (*senderrmqworker.Worker, error) {
	producer := rmqconsumer.NewConsumer(
		"",
		config.RMQ.Uri,
		config.RMQ.Exchange,
		config.RMQ.ExchangeType,
		config.RMQ.Queue,
		config.RMQ.Key,
		10*time.Second)

	return senderrmqworker.NewWorker(logger, *producer), nil
}
