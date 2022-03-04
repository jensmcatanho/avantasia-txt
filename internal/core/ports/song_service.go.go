package ports

import (
	"context"

	"github.com/jensmcatanho/avantasia-txt/internal/core/domain"
)

type SongService interface {
	GetRandomSong(ctx context.Context) (*domain.Song, error)
	GetSongByID(ctx context.Context, id int) (*domain.Song, error)
}
