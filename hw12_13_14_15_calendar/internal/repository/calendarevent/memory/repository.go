package memoryrepository

import (
	"strconv"

	domain "github.com/fixme_my_friend/hw12_13_14_15_calendar/domain/calendarevent"
	memorystorage "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage/memory"
)

type CalendarEventRepository struct {
	storage *memorystorage.Storage
}

func NewCalendarEventRepository(storage *memorystorage.Storage) *CalendarEventRepository {
	return &CalendarEventRepository{
		storage: storage,
	}
}

func (repository *CalendarEventRepository) GetEvents(userID int) ([]domain.CalendarEventEntity, error) {
	entity, err := repository.storage.Find(strconv.Itoa(userID))

	return []domain.CalendarEventEntity{entity}, err
}

func (repository *CalendarEventRepository) Insert(entities []domain.CalendarEventEntity) error {
	for _, entity := range entities {
		err := repository.storage.Insert(strconv.Itoa(entity.UserID), entity)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repository *CalendarEventRepository) Update(entity domain.CalendarEventEntity) error {
	err := repository.storage.Update(strconv.Itoa(entity.UserID), entity)
	if err != nil {
		return err
	}

	return nil
}

func (repository *CalendarEventRepository) Delete(userID int) error {
	err := repository.storage.Delete(strconv.Itoa(userID))
	if err != nil {
		return err
	}

	return nil
}
