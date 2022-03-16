package app

import (
	domain "github.com/fixme_my_friend/hw12_13_14_15_calendar/domain/calendarevent"
	"time"
)

func (a *App) GetEvents(userId int) ([]domain.CalendarEventEntity, error) {
	events, err := a.calendarEventRepository.GetEvents(userId)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (a *App) CreateEvent(title string, dateTimeEvent time.Time, durationEvent string, userId int) error {
	event := &domain.CalendarEventEntity{
		UserId:        userId,
		Title:         title,
		DateTimeEvent: dateTimeEvent,
		DurationEvent: durationEvent,
	}

	return a.calendarEventRepository.Insert([]domain.CalendarEventEntity{*event})
}

func (a *App) UpdateEvent(userId int, title string) error {
	event := &domain.CalendarEventEntity{
		UserId: userId,
		Title:  title,
	}

	return a.calendarEventRepository.Update(*event)
}

func (a *App) DeleteEvent(userId int) error {
	return a.calendarEventRepository.Delete(userId)
}
