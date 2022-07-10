package searchapi

import (
	"context"
)

// Music ...
type Music struct {
	ID    int64
	Title string
}

//go:generate mockgen -package mockservice -source ./musicbox.go -destination ./mockservice/musicbox.go

// Service ...
type Service interface {
	GetMusicByTitle(ctx context.Context, title string) ([]Music, error)
}
