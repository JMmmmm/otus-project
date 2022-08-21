package sqlrepository

import (
	"context"
	"fmt"

	domain "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/domain/calendarevent"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/logger"

	// nolint
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type CalendarEventRepository struct {
	Logger logger.Logger
	DB     *sqlx.DB
	Ctx    *context.Context
}

func NewCalendarEventRepository(
	logger logger.Logger,
	ctx context.Context,
	dsn string) (*CalendarEventRepository, error) {
	repository := &CalendarEventRepository{
		Logger: logger,
	}

	err := repository.connect(ctx, dsn)

	return repository, err
}

func (repository *CalendarEventRepository) GetEvents(userID int) ([]domain.CalendarEventEntity, error) {
	sql := `
		SELECT 
			id,
		    user_id,
		    title,
		    datetime_event,
		    duration_event,
		    coalesce(description, '') as description,
		    coalesce(notification_interval, '0') as notification_interval
		FROM public.calendar_event where user_id = :userID
	`

	rows, err := repository.DB.NamedQueryContext(*repository.Ctx, sql, map[string]interface{}{
		"userID": userID,
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

	var events []domain.CalendarEventEntity
	for rows.Next() {
		var event domain.CalendarEventEntity
		err := rows.StructScan(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, rows.Err()
}

func (repository *CalendarEventRepository) InsertEntities(entities []domain.CalendarEventEntity) error {
	_, err := repository.DB.NamedExec(`
		INSERT INTO public.calendar_event (user_id, title, datetime_event, duration_event)
        VALUES (:user_id, :title, :datetime_event, :duration_event)
	`, entities)

	return err
}

func (repository *CalendarEventRepository) Update(entity domain.CalendarEventEntity) error {
	_, err := repository.DB.NamedExec(`
		UPDATE public.calendar_event set title = :title, datetime_event = :datetime_event
        where id = :id
	`, entity)

	return err
}

func (repository *CalendarEventRepository) Delete(id string) error {
	_, err := repository.DB.Exec(`DELETE from public.calendar_event where id = $1`, id)

	return err
}

func (repository *CalendarEventRepository) connect(ctx context.Context, dsn string) (err error) {
	repository.Ctx = &ctx

	repository.DB, err = sqlx.Open("pgx", dsn)
	if err != nil {
		return fmt.Errorf("cannot open pgx driver: %w", err)
	}

	return repository.DB.PingContext(ctx)
}

func (repository *CalendarEventRepository) Close(ctx context.Context) error {
	return repository.DB.Close()
}
