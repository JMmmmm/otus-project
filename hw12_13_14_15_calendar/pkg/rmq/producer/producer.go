package rmqproducer

import (
	"fmt"

	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/logger"
	"github.com/streadway/amqp"
)

type Producer struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	logger     logger.Logger
	queue      amqp.Queue
}

func NewProducer(logg logger.Logger) *Producer {
	return &Producer{
		logger: logg,
	}
}

func (producer *Producer) Connect(amqpURI string, exchange string, exchangeType string, queueName string) (err error) {
	producer.logger.Info(fmt.Sprintf("dialing %q", amqpURI))
	producer.connection, err = amqp.Dial(amqpURI)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	producer.logger.Info("got Connection, getting Channel")
	producer.channel, err = producer.connection.Channel()
	if err != nil {
		return fmt.Errorf("channel: %w", err)
	}

	producer.logger.Info(fmt.Sprintf("got Channel, declaring %q Exchange (%q)", exchangeType, exchange))
	if err := producer.channel.ExchangeDeclare(
		exchange,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		return fmt.Errorf("exchange Declare: %w", err)
	}

	producer.logger.Info("declared Exchange")

	producer.queue, err = producer.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	return err
}

func (producer *Producer) Publish(exchange string, routingKey string, body string, reliable bool) (err error) {
	// Reliable publisher confirms require confirm.select support from the
	// connection.
	if reliable {
		producer.logger.Info("enabling publishing confirms.")
		if err := producer.channel.Confirm(false); err != nil {
			return fmt.Errorf("channel could not be put into confirm mode: %w", err)
		}

		confirms := producer.channel.NotifyPublish(make(chan amqp.Confirmation, 1))

		defer producer.confirmOne(confirms)
	}

	headers := make(amqp.Table)
	headers["ss"] = "ee"

	if err = producer.channel.Publish(
		exchange,
		producer.queue.Name,
		false,
		false,
		amqp.Publishing{
			Headers:         headers,
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(body),
			DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
			Priority:        0,
		},
	); err != nil {
		return fmt.Errorf("exchange Publish: %w", err)
	}

	return nil
}

func (producer *Producer) Close() error {
	return producer.connection.Close()
}

// One would typically keep a channel of publishings, a sequence number, and a
// set of unacknowledged sequence numbers and loop until the publishing channel
// is closed.
func (producer *Producer) confirmOne(confirms <-chan amqp.Confirmation) {
	producer.logger.Info("waiting for confirmation of one publishing")

	if confirmed := <-confirms; confirmed.Ack {
		producer.logger.Info(fmt.Sprintf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag))
	} else {
		producer.logger.Info(fmt.Sprintf("failed delivery of delivery tag: %d", confirmed.DeliveryTag))
	}
}
