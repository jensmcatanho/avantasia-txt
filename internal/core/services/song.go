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
	song, _, err := s.songRepository.GetSongByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (s *songService) GetSongByName(ctx context.Context, name string) (*domain.Song, error) {
	song, _, err := s.songRepository.GetSongByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (s *songService) UpdateSongLyric(ctx context.Context, name string, lyricID string, newLyric string) error {
	parsedLyricID, err := strconv.Atoi(lyricID)
	if err != nil {
		return err
	}

	song, referenceID, err := s.songRepository.GetSongByName(ctx, name)
	if err != nil {
		return err
	}

	song.Lyrics[parsedLyricID] = newLyric
	err = s.songRepository.UpdateSong(ctx, referenceID, song)
	if err != nil {
		return err
	}

	return nil
}
