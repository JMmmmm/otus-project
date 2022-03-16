package memoryrepository

import (
	domain "github.com/fixme_my_friend/hw12_13_14_15_calendar/domain/calendarevent"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
	memorystorage "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage/memory"
	"strconv"
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
	entity, err := repository.storage.Find(strconv.Itoa(userId))

	return []domain.CalendarEventEntity{entity}, err
}

func (repository *CalendarEventRepository) Insert(entities []domain.CalendarEventEntity) error {
	for _, entity := range entities {
		err := repository.storage.Insert(strconv.Itoa(entity.UserId), &entity)
		if err != nil {
			return err
		}
	}

	return nil
}
func (repository *CalendarEventRepository) Update(entity domain.CalendarEventEntity) error {
	err := repository.storage.Insert(strconv.Itoa(entity.UserId), &entity)
	if err != nil {
		return err
	}

	return nil
}
func (repository *CalendarEventRepository) Delete(userId int) error {
	err := repository.storage.Delete(strconv.Itoa(userId))
	if err != nil {
		return err
	}

	return nil
}
