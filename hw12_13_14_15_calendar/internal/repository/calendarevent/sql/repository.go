package sqlrepository

import (
	"fmt"

	domain "github.com/fixme_my_friend/hw12_13_14_15_calendar/domain/calendarevent"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
	sqlstorage "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage/sql"
)

type CalendarEventRepository struct {
	storage *sqlstorage.Storage
	Logger  logger.Logger
}

func NewCalendarEventRepository(storage *sqlstorage.Storage, logger logger.Logger) *CalendarEventRepository {
	return &CalendarEventRepository{
		storage: storage,
		Logger:  logger,
	}
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

	rows, err := repository.storage.DB.NamedQueryContext(*repository.storage.Ctx, sql, map[string]interface{}{
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

func (repository *CalendarEventRepository) Insert(entities []domain.CalendarEventEntity) error {
	_, err := repository.storage.DB.NamedExec(`
		INSERT INTO public.calendar_event (user_id, title, datetime_event, duration_event)
        VALUES (:user_id, :title, :datetime_event, :duration_event)
	`, entities)

	return err
}

func (repository *CalendarEventRepository) Update(entity domain.CalendarEventEntity) error {
	_, err := repository.storage.DB.NamedExec(`
		UPDATE public.calendar_event set title = :title, datetime_event = :datetime_event
        where user_id = :user_id
	`, entity)

	return err
}

func (repository *CalendarEventRepository) Delete(userID int) error {
	_, err := repository.storage.DB.Exec(`DELETE from public.calendar_event where user_id = $1`, userID)

	return err
}
