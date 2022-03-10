package sqlrepository

import (
	"fmt"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/domain"
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

func (repository *CalendarEventRepository) GetEvents(userId int) ([]domain.CalendarEventEntity, error) {
	sql := `
		SELECT * FROM public.calendar_event where user_id = :userId
	`

	rows, err := repository.storage.Db.NamedQueryContext(*repository.storage.Ctx, sql, map[string]interface{}{
		"userId": userId,
		//"datetime_event": "2019-12-31",
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

func (repository *CalendarEventRepository) Insert(entities []domain.CalendarEvent) error {
	_, err := repository.storage.Db.NamedExec(`INSERT INTO public.calendar_event (user_id, title, datetime_event, duration_event, description, notification_interval)
        VALUES (:user_id, :title, :datetime_event, :duration_event, :description, :notification_interval)`, entities)

	return err
}

func (repository *CalendarEventRepository) Update(entity domain.CalendarEvent) error {
	_, err := repository.storage.Db.NamedExec(`UPDATE public.calendar_event set title = :title, datetime_event = :datetime_event
        where user_id = :user_id`, entity)

	return err
}

func (repository *CalendarEventRepository) Delete(userId int) error {
	_, err := repository.storage.Db.Exec(`DELETE from public.calendar_event where user_id = $1`, userId)

	return err
}
