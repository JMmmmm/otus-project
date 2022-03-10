package memorystorage

import "sync"

type Storage struct {
	internalStorage map[string]interface{}
	mu              sync.RWMutex
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Get() {
}

// TODO
