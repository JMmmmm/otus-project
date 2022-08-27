package sender

import (
	"context"
	"fmt"
	sqlrepository "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/repository/notification/sql"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/logger"
	"time"

	senderrmqworker "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/app/sender/rmqworker"
	rmqconsumer "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/rmq/consumer"
)

func CreateWorker(ctx context.Context, config Config, logger logger.Logger) (*senderrmqworker.Worker, error) {
	repository, err := sqlrepository.NewNotificationRepository(ctx, logger, config.PSQL.DSN)
	if err != nil {
		return nil, fmt.Errorf("can not create repository: %w", err)
	}

	producer := rmqconsumer.NewConsumer(
		"",
		config.RMQ.URI,
		config.RMQ.Exchange,
		config.RMQ.ExchangeType,
		config.RMQ.Queue,
		config.RMQ.Key,
		10*time.Second)

	return senderrmqworker.NewWorker(logger, *producer, repository), nil
}
