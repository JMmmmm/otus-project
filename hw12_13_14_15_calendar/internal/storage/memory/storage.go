package memorystorage

import (
	"errors"
	domain "github.com/fixme_my_friend/hw12_13_14_15_calendar/domain/calendarevent"
	"sync"
)

type Storage struct {
	mu    sync.RWMutex
	cache map[string]domain.CalendarEventEntity
}

func New() *Storage {
	return &Storage{
		cache: make(map[string]domain.CalendarEventEntity),
	}
}

func (s *Storage) Find(key string) (domain.CalendarEventEntity, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if val, found := s.cache[key]; found {
		return val, nil
	}

	return domain.CalendarEventEntity{}, errors.New("not found")
}

func (s *Storage) Get(offset, limit uint) []domain.CalendarEventEntity {
	list := make([]domain.CalendarEventEntity, 0, limit)
	s.mu.Lock()
	defer s.mu.Unlock()

	i := uint(0)
	for _, val := range s.cache {
		if i < offset {
			i++
			continue
		}
		list = append(list, val)
	}

	return list
}

func (s *Storage) Insert(key string, value *domain.CalendarEventEntity) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache[key] = *value
	return nil
}

func (s *Storage) Update(key string, value *domain.CalendarEventEntity) error {
	return s.Insert(key, value)
}

func (s *Storage) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.cache, key)
	return nil
}
