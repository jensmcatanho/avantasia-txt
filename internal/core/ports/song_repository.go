package ports

import (
	"context"

	"github.com/jensmcatanho/avantasia-txt/internal/core/domain"
)

type SongRepository interface {
	GetSongByID(ctx context.Context, id int) (*domain.Song, string, error)
	GetSongByName(ctx context.Context, name string) (*domain.Song, string, error)
	UpdateSong(ctx context.Context, referenceID string, song *domain.Song) error
}
