package services

import (
	"fmt"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/jensmcatanho/avantasia-txt/internal/core/domain"
)

const replyDelay = 100 * time.Millisecond

type twitterService struct {
	client *twitter.Client
}

func NewTwitterService(credentials *domain.TwitterCredentials) *twitterService {
	config := oauth1.NewConfig(credentials.ConsumerKey, credentials.ConsumerSecret)
	token := oauth1.NewToken(credentials.AccessToken, credentials.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	return &twitterService{
		client: client,
	}
}

func (t *twitterService) Tweet(song *domain.Song) error {
	tweet, _, err := t.client.Statuses.Update(song.GetLyric(), nil)
	if err != nil {
		return err
	}

	time.Sleep(replyDelay)

	_, _, err = t.client.Statuses.Update(fmt.Sprintf("Song: %s\nAlbum: %s", song.Name, song.Album), &twitter.StatusUpdateParams{
		InReplyToStatusID: tweet.ID,
	})
	if err != nil {
		return err
	}

	return nil
}
