package service

import (
    "7-url-shortener/internal/storage"
    "7-url-shortener/internal/util"
)

type ShortenService struct {
    storage *storage.Storage
}

func NewShortenService() *ShortenService {
    return &ShortenService{
        storage: storage.NewStorage(),
    }
}

func (s *ShortenService) Shorten(originalURL string) string {
    shortCode := util.GenerateShortCode()

    for {
        if _, exists := s.storage.Get(shortCode); !exists {
            break
        }
        shortCode = util.GenerateShortCode()
    }

    s.storage.Save(shortCode, originalURL)
    return shortCode
}

func (s *ShortenService) Lookup(shortCode string) (string, bool) {
    return s.storage.Get(shortCode)
}
