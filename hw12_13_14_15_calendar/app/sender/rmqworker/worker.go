package senderrmqworker

import (
	"context"
	"log"

	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/logger"
	rmqconsumer "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/rmq/consumer"
	"github.com/streadway/amqp"
)

type Worker struct {
	consumer rmqconsumer.Consumer
	logger   logger.Logger
}

func NewWorker(logg logger.Logger, consumer rmqconsumer.Consumer) *Worker {
	return &Worker{
		logger:   logg,
		consumer: consumer,
	}
}

func (worker Worker) Execute(ctx context.Context, threads int) error {
	log.Print("Fixed Rate of 5 seconds")

	return worker.consumer.Handle(ctx, work, threads)
}

func work(ctx context.Context, deliveries <-chan amqp.Delivery) {
	for d := range deliveries {
		log.Printf(
			"got %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)
		err := d.Ack(false)
		if err != nil {
			log.Printf("got err on delivery: %v", err)
		}
	}
}
