package memoryrepository

import (
	"errors"
	"fmt"
	"strconv"
	"sync"

	domain "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/domain/calendarevent"
)

type CalendarEventRepository struct {
	mu             sync.RWMutex
	cache          map[string]domain.CalendarEventEntity
	cacheByUserIds map[string]map[string]domain.CalendarEventEntity
}

func NewCalendarEventRepository() *CalendarEventRepository {
	return &CalendarEventRepository{
		cache:          make(map[string]domain.CalendarEventEntity),
		cacheByUserIds: make(map[string]map[string]domain.CalendarEventEntity),
	}
}

func (repository *CalendarEventRepository) GetEvents(userID int) ([]domain.CalendarEventEntity, error) {
	repository.mu.Lock()
	defer repository.mu.Unlock()
	if values, found := repository.cacheByUserIds[strconv.Itoa(userID)]; found {
		result := make([]domain.CalendarEventEntity, len(values))
		i := 0
		for _, entity := range values {
			result[i] = entity
			i++
		}

		return result, nil
	}

	return []domain.CalendarEventEntity{}, errors.New("not found")
}

func (repository *CalendarEventRepository) InsertEntities(entities []domain.CalendarEventEntity) error {
	var err error
	for _, entity := range entities {
		userID := strconv.Itoa(entity.UserID)

		repository.mu.Lock()

		if _, found := repository.cache[entity.ID]; found {
			err = fmt.Errorf("not found in cache: %w", err)
		}

		existUserEntities, found := repository.cacheByUserIds[userID]
		if found {
			err = fmt.Errorf("not found in cache indexed by users ids: %w", err)
		}

		repository.cache[entity.ID] = entity

		if len(existUserEntities) == 0 {
			entitiesMap := make(map[string]domain.CalendarEventEntity)
			entitiesMap[entity.ID] = entity
			repository.cacheByUserIds[userID] = entitiesMap
		} else {
			repository.cacheByUserIds[userID][entity.ID] = entity
		}

		repository.mu.Unlock()
	}

	return err
}

func (repository *CalendarEventRepository) Update(entity domain.CalendarEventEntity) error {
	userID := strconv.Itoa(entity.UserID)
	repository.mu.Lock()
	defer repository.mu.Unlock()

	if _, found := repository.cache[entity.ID]; !found {
		return errors.New("not found in cache")
	}
	if _, found := repository.cacheByUserIds[userID]; !found {
		return errors.New("not found in cache indexed by users ids")
	}

	repository.cache[entity.ID] = entity
	repository.cacheByUserIds[userID][entity.ID] = entity

	return nil
}

func (repository *CalendarEventRepository) Delete(id string) error {
	repository.mu.Lock()
	defer repository.mu.Unlock()

	entity, found := repository.cache[id]
	if !found {
		return errors.New("not found in cache")
	}

	userID := strconv.Itoa(entity.UserID)
	userEntities, found := repository.cacheByUserIds[userID]
	if !found {
		return errors.New("not found in cache indexed by user id")
	}

	if _, found := userEntities[id]; !found {
		return errors.New("not found id in cache indexed by users ids")
	}

	delete(repository.cache, id)
	delete(userEntities, id)
	if len(userEntities) == 0 {
		delete(repository.cacheByUserIds, userID)
	} else {
		repository.cacheByUserIds[userID] = userEntities
	}

	return nil
}
