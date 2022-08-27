package queue

type Producer interface {
	Connect(amqpURI string, exchange string, exchangeType string, queueName string) (err error)
	Publish(exchange string, routingKey string, body string, reliable bool) (err error)
	Close() error
}
