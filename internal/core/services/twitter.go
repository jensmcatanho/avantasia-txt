package services

import (
	"fmt"
	"math/rand"
	"strconv"
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

func (t *twitterService) TweetLyric(song *domain.Song, lyricID string) error {
	lyric, err := strconv.Atoi(lyricID)
	if err != nil {
		return err
	}

	return t.tweetLyric(song, lyric)
}

func (t *twitterService) TweetRandomLyric(song *domain.Song) error {
	return t.tweetLyric(song, rand.Intn(len(song.Lyrics)))
}

func (t *twitterService) tweetLyric(song *domain.Song, lyricID int) error {
	tweet, _, err := t.client.Statuses.Update(song.Lyrics[lyricID], nil)
	if err != nil {
		return err
	}

	err = t.reply(tweet.ID, song)
	if err != nil {
		return err
	}

	return nil
}

func (t *twitterService) reply(tweetID int64, song *domain.Song) error {
	time.Sleep(replyDelay)

	_, _, err := t.client.Statuses.Update(fmt.Sprintf("Song: %s\nAlbum: %s", song.Name, song.Album), &twitter.StatusUpdateParams{
		InReplyToStatusID: tweetID,
	})
	if err != nil {
		return err
	}

	return nil
}
