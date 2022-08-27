package senderrmqworker

import (
	"context"
	"encoding/json"
	"log"

	domain "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/domain/notification"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/logger"
	rmqconsumer "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/rmq/consumer"
	"github.com/streadway/amqp"
)

type Worker struct {
	consumer   rmqconsumer.Consumer
	logger     logger.Logger
	repository domain.NotificationRepository
}

func NewWorker(logg logger.Logger, consumer rmqconsumer.Consumer, repository domain.NotificationRepository) *Worker {
	return &Worker{
		logger:     logg,
		consumer:   consumer,
		repository: repository,
	}
}

func (worker Worker) Execute(ctx context.Context, threads int) error {
	return worker.consumer.Handle(ctx, worker.work, threads)
}

func (worker Worker) work(ctx context.Context, deliveries <-chan amqp.Delivery) {
	for d := range deliveries {
		log.Printf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)

		var entity domain.NotificationEntity
		err := json.Unmarshal(d.Body, &entity)
		if err != nil {
			log.Printf("got err on unmarshal: %v", err)
			return
		}

		err = worker.repository.Update(entity)
		if err != nil {
			log.Printf("can not update notification: %v", err)
		}

		err = d.Ack(false)
		if err != nil {
			log.Printf("got err on delivery: %v", err)
		}
	}
}
