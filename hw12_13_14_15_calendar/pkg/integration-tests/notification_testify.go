package integration_tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/logger"
	rmqproducer "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/rmq/producer"
	"github.com/cucumber/godog"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

var (
	dbDSN           = os.Getenv("TESTS_DB_DSN")
	amqpUri         = os.Getenv("TESTS_AMQP_URI")
	amqpExchange    = "test-exchange-2"
	amqpEchangeType = "direct"
	amqpKey         = "test-key"
	amqpReliable    = true
	amqpQueue       = "test-2"
)

func init() {
	if amqpUri == "" {
		amqpUri = "amqp://guest:guest@localhost:5672/"
	}
	if dbDSN == "" {
		dbDSN = "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"
	}
}

type NotificationEntity struct {
	ID     string
	UserID int `db:"user_id"`
	Title  string
}

type notifyTest struct {
	producer *rmqproducer.Producer
	db       *sqlx.DB

	notificationId string
}

func (test *notifyTest) start(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
	logg, err := logger.NewAppLogger("INFO", "stdout")
	if err != nil {
		return nil, fmt.Errorf("can not create logger: %w", err)
	}

	test.producer = rmqproducer.NewProducer(logg)
	err = test.producer.Connect(amqpUri, amqpExchange, amqpEchangeType, amqpQueue)
	if err != nil {
		return ctx, err
	}
	test.db, err = sqlx.Open("pgx", dbDSN)
	if err != nil {
		return ctx, fmt.Errorf("cannot open pgx driver: %w", err)
	}
	err = test.db.PingContext(ctx)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func (test *notifyTest) stop(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
	err = test.producer.Close()
	if err != nil {
		return ctx, err
	}
	err = test.db.Close()
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func (test *notifyTest) theNotificationShouldBe(title string) error {
	time.Sleep(5 * time.Second)
	sql := `
		SELECT 
			id,
		    user_id,
		    title
		FROM public.calendar_event where id=$1
	`

	var notification NotificationEntity
	err := test.db.Get(&notification, sql, test.notificationId)
	if err != nil {
		return err
	}

	if notification.Title != title {
		return fmt.Errorf("unexpected title: %s != %s", notification.Title, title)
	}
	return nil
}

func (test *notifyTest) iPublishNotificationWithData(data string) (err error) {
	var notification NotificationEntity
	err = json.Unmarshal([]byte(data), &notification)
	if err != nil {
		return err
	}

	test.notificationId = notification.ID

	return test.producer.Publish(amqpExchange, amqpKey, data, amqpReliable)
}

func NotificationFeatureContext(s *godog.ScenarioContext) {
	test := new(notifyTest)

	s.Before(test.start)

	s.Step(`^I publish notification with data:$`, test.iPublishNotificationWithData)
	s.Step(`^The notification should be "([^"]*)"$`, test.theNotificationShouldBe)

	s.After(test.stop)

}
