package handlers

import "github.com/garrettladley/roundest/internal/storage"

type Service struct {
	store storage.Storage
}

func NewService(store storage.Storage) *Service {
	return &Service{store: store}
}
