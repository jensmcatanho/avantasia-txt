package services

import (
	"context"
	"math/rand"
	"os"
	"strconv"

	"github.com/jensmcatanho/avantasia-txt/internal/core/domain"
	"github.com/jensmcatanho/avantasia-txt/internal/core/ports"
)

type songService struct {
	songRepository ports.SongRepository
}

func NewSongService(songRepository ports.SongRepository) *songService {
	return &songService{
		songRepository: songRepository,
	}
}

func (s *songService) GetRandomSong(ctx context.Context) (*domain.Song, error) {
	totalSongs, err := strconv.Atoi(os.Getenv("TOTAL_SONGS"))
	if err != nil {
		return nil, err
	}

	songID := rand.Intn(totalSongs) + 1
	return s.getSongByID(ctx, songID)
}

func (s *songService) GetSongByID(ctx context.Context, id int) (*domain.Song, error) {
	return s.getSongByID(ctx, id)
}

func (s *songService) getSongByID(ctx context.Context, id int) (*domain.Song, error) {
	song, err := s.songRepository.GetSongByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return song, nil
}
