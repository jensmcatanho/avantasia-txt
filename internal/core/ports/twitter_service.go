package ports

import "github.com/jensmcatanho/avantasia-txt/internal/core/domain"

type TwitterService interface {
	TweetLyric(song *domain.Song, lyricID string) error
	TweetRandomLyric(song *domain.Song) error
}
