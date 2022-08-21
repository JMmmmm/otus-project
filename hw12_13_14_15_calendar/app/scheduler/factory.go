package scheduler

import (
	"context"
	"fmt"

	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/logger"
	sqlrepository "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/repository/notification/sql"
	rmqproducer "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/rmq/producer"
)

func CreateWorker(ctx context.Context, config Config, logger logger.Logger) (*Worker, error) {
	repository, err := sqlrepository.NewNotificationRepository(logger, ctx, config.PSQL.DSN)
	if err != nil {
		return nil, fmt.Errorf("can not create repository: %v", err)
	}

	producer := rmqproducer.NewProducer(logger)

	return NewWorker(logger, repository, producer), nil
}
