package memoryrepository

import (
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/domain"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
	memorystorage "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage/memory"
)

type CalendarEventRepository struct {
	storage *memorystorage.Storage
	logger  logger.Logger
}

func NewCalendarEventRepository(storage *memorystorage.Storage, logger logger.Logger) *CalendarEventRepository {
	return &CalendarEventRepository{
		storage: storage,
	}
}

func (repository *CalendarEventRepository) GetEvents(userId int) ([]domain.CalendarEventEntity, error) {
	var entities []domain.CalendarEventEntity

	return entities, nil
}

func (repository *CalendarEventRepository) Insert(entities []domain.CalendarEvent) error {
	return nil
}
func (repository *CalendarEventRepository) Update(entity domain.CalendarEvent) error {
	return nil
}
func (repository *CalendarEventRepository) Delete(userId int) error {
	return nil
}
