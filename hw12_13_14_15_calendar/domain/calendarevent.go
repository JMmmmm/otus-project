package domain

import (
	"time"
)

type CalendarEvent struct {
	UserId               int `db:"user_id"`
	Title                string
	DateTimeEvent        time.Time `db:"datetime_event"`
	DurationEvent        string    `db:"duration_event"`
	Description          string
	NotificationInterval string `db:"notification_interval"`
}

type CalendarEventEntity struct {
	Id int
	CalendarEvent
}

func NewCalendarEvent(
	userId int,
	title string,
	dateTimeEvent time.Time,
	durationEvent string,
	description string,
	notificationInterval string) *CalendarEvent {
	calendarEvent := CalendarEvent{
		UserId:               userId,
		Title:                title,
		DateTimeEvent:        dateTimeEvent,
		DurationEvent:        durationEvent,
		Description:          description,
		NotificationInterval: notificationInterval,
	}

	return &calendarEvent
}

type CalendarEventRepository interface {
	GetEvents(userId int) ([]CalendarEventEntity, error)
	Insert(entities []CalendarEvent) error
	Update(entity CalendarEvent) error
	Delete(userId int) error
}
