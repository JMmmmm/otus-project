package domain

import "time"

type NotificationEntity struct {
	ID            string
	UserID        int `db:"user_id"`
	Title         string
	DateTimeEvent time.Time `db:"datetime_event"`
}
