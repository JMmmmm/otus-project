package calendar

import (
	"time"

	domain "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/domain/calendarevent"
)

func (a *App) GetEvents(userID int) ([]domain.CalendarEventEntity, error) {
	events, err := a.calendarEventRepository.GetEvents(userID)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (a *App) CreateEvent(title string, dateTimeEvent time.Time, durationEvent string, userID int) error {
	event := domain.CalendarEventEntity{
		UserID:        userID,
		Title:         title,
		DateTimeEvent: dateTimeEvent,
		DurationEvent: durationEvent,
	}

	return a.calendarEventRepository.InsertEntities([]domain.CalendarEventEntity{event})
}

func (a *App) UpdateEvent(id string, title string) error {
	event := domain.CalendarEventEntity{
		ID:    id,
		Title: title,
	}

	return a.calendarEventRepository.Update(event)
}

func (a *App) DeleteEvent(id string) error {
	return a.calendarEventRepository.Delete(id)
}
