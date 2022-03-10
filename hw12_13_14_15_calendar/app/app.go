package app

import (
	"context"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/domain"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
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

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

func (a *App) GetEvents(userId int) ([]domain.CalendarEventEntity, error) {
	events, err := a.calendarEventRepository.GetEvents(userId)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (a *App) InsertEvent(entities []domain.CalendarEvent) error {
	return a.calendarEventRepository.Insert(entities)
}

func (a *App) UpdateEvent(entity domain.CalendarEvent) error {
	return a.calendarEventRepository.Update(entity)
}

func (a *App) DeleteEvent(userId int) error {
	return a.calendarEventRepository.Delete(userId)
}
