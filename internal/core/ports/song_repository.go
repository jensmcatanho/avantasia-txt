package ports

import (
	"context"

	"github.com/jensmcatanho/avantasia-txt/internal/core/domain"
)

type SongRepository interface {
	GetSongByID(ctx context.Context, id int) (*domain.Song, error)
	GetSongByName(ctx context.Context, name string) (*domain.Song, error)
}
