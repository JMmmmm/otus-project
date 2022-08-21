package domain

import "time"

type NotificationRepository interface {
	GetNotifications(timeFrom time.Time, timeTo time.Time) ([]NotificationEntity, error)
}
