package storage

import "sync"

type Storage struct {
	mu   sync.RWMutex
	urls map[string]string
}

func NewStore() *Storage {
	return &Storage{
		urls: make(map[string]string),
	}
}

func (s *Storage) Save(shortCode, originalUrl string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.urls[shortCode] = originalUrl
}

func (s *Storage) Get(shortCode string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, ok := s.urls[shortCode]
	return url, ok
}
