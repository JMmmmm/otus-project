package scheduler

import (
	"encoding/json"
	"fmt"
	"time"

	domain "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/domain/notification"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/logger"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/queue"
)

type Worker struct {
	repository domain.NotificationRepository
	producer   queue.Producer
	logger     logger.Logger
}

func NewWorker(logg logger.Logger, repository domain.NotificationRepository, producer queue.Producer) *Worker {
	return &Worker{
		logger:     logg,
		repository: repository,
		producer:   producer,
	}
}

func (worker Worker) Execute(
	uri string,
	exchangeType string,
	queueName string,
	exchange string,
	routingKey string,
	reliable bool,
	timeFrom time.Time,
	timeTo time.Time) {
	err := worker.producer.Connect(uri, exchange, exchangeType, queueName)
	if err != nil {
		worker.logger.Error(fmt.Sprintf("can not connect to rmq: %v", err))
		return
	}
	notifications, err := worker.repository.GetNotifications(timeFrom, timeTo)
	if err != nil {
		worker.logger.Error(fmt.Sprintf("Can not get notifications: %v", err))
		return
	}

	for _, notification := range notifications {
		body, err := json.Marshal(notification)
		if err != nil {
			worker.logger.Error(fmt.Sprintf("Can not json encode: %v", err))
			continue
		}

		err = worker.producer.Publish(exchange, routingKey, string(body), reliable)
		if err != nil {
			worker.logger.Error(fmt.Sprintf("Can not publish message: %v", err))
			continue
		}
	}

	err = worker.producer.Close()
	if err != nil {
		worker.logger.Error(fmt.Sprintf("can not close connection rmq: %v", err))
		return
	}
}
