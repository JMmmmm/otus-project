package sqlrepository

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/stdlib" //nolint
	"github.com/jmoiron/sqlx"
	"time"

	domain "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/domain/notification"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/logger"
)

type NotificationRepository struct {
	Logger logger.Logger
	DB     *sqlx.DB
	Ctx    *context.Context
}

func NewNotificationRepository(logger logger.Logger, ctx context.Context, dsn string) (*NotificationRepository, error) {
	repository := &NotificationRepository{
		Logger: logger,
	}

	err := repository.connect(ctx, dsn)

	return repository, err
}

func (repository *NotificationRepository) GetNotifications(timeFrom time.Time, timeTo time.Time) ([]domain.NotificationEntity, error) {
	sql := `
		SELECT 
			id,
		    user_id,
		    title,
		    datetime_event
		FROM public.calendar_event where datetime_event > :timeFrom and datetime_event < :timeTo
	`

	rows, err := repository.DB.NamedQueryContext(*repository.Ctx, sql, map[string]interface{}{
		"timeFrom": timeFrom,
		"timeTo":   timeTo,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot select: %w", err)
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			repository.Logger.Error(fmt.Sprintf("cannot select: %v", err))
		}
	}()

	var events []domain.NotificationEntity
	for rows.Next() {
		var event domain.NotificationEntity
		err := rows.StructScan(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, rows.Err()
}

func (repository *NotificationRepository) connect(ctx context.Context, dsn string) (err error) {
	repository.Ctx = &ctx

	repository.DB, err = sqlx.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("cannot open pgx driver: %w", err)
	}

	return repository.DB.PingContext(ctx)
}

func (repository *NotificationRepository) Close(ctx context.Context) error {
	return repository.DB.Close()
}
