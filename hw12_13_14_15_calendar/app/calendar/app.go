package calendar

import (
	"time"

	domain "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/domain/calendarevent"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/logger"
)

type Application interface {
	GetEvents(userID int) ([]domain.CalendarEventEntity, error)
	CreateEvent(title string, dateTimeEvent time.Time, DurationEvent string, userID int) error
	UpdateEvent(id string, title string) error
	DeleteEvent(id string) error
}

type App struct {
	logger                  logger.Logger
	calendarEventRepository domain.CalendarEventRepository
}

func New(logger logger.Logger, calendarEventRepository domain.CalendarEventRepository) *App {
	return &App{logger: logger, calendarEventRepository: calendarEventRepository}
}
