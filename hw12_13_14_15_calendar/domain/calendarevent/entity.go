package domain

import "time"

type CalendarEventEntity struct {
	ID                   string
	UserId               int `db:"user_id"`
	Title                string
	DateTimeEvent        time.Time `db:"datetime_event"`
	DurationEvent        string    `db:"duration_event"`
	Description          string
	NotificationInterval string `db:"notification_interval"`
}
