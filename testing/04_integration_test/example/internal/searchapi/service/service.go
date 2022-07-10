package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/steveanlorn/learning-go/testing/04_integration_test/example/internal/searchapi"
)

// Music ...
type Music struct {
	ID    int64
	Title string
}

//go:generate mockgen -package service -source ./service.go -destination ./storemock_test.go

// Store ...
type Store interface {
	GetMusicByTitle(ctx context.Context, title string) ([]Music, error)
}

// Service ...
type Service struct {
	store Store
}

// NewService ...
func NewService(store Store) *Service {
	service := Service{
		store: store,
	}

	return &service
}

// GetMusicByTitle ...
func (s *Service) GetMusicByTitle(ctx context.Context, title string) ([]searchapi.Music, error) {
	if title == "" {
		return nil, errors.New("title is reqiured")
	}

	musics, err := s.store.GetMusicByTitle(ctx, title)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("GetMusicByTitle(ctx, %s) -> s.store.GetMusicByTitle(ctx, %s)", title, title))
	}

	result := make([]searchapi.Music, len(musics))
	for i, music := range musics {
		result[i] = searchapi.Music{
			ID:    music.ID,
			Title: music.Title,
		}
	}

	return result, nil
}
