package ports

import "github.com/jensmcatanho/avantasia-txt/internal/core/domain"

type TwitterService interface {
	Tweet(song *domain.Song) error
}
