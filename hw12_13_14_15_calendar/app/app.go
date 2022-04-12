package app

import (
	domain "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/domain/calendarevent"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/logger"
)

type App struct {
	logger                  logger.Logger
	calendarEventRepository domain.CalendarEventRepository
}

type Storage interface {
	Get()
}

func New(logger logger.Logger, calendarEventRepository domain.CalendarEventRepository) *App {
	return &App{logger: logger, calendarEventRepository: calendarEventRepository}
}
