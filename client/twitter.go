package client

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/jensmcatanho/avantasia-txt/models"
)

type TwitterClient struct {
	client *twitter.Client
}

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func NewTwitterClient(credentials Credentials) *TwitterClient {
	config := oauth1.NewConfig(credentials.ConsumerKey, credentials.ConsumerSecret)
	token := oauth1.NewToken(credentials.AccessToken, credentials.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	return &TwitterClient{
		client: client,
	}
}

func (t *TwitterClient) Tweet(song *models.Song) error {
	tweet, _, err := t.client.Statuses.Update(song.GetLyric(), nil)
	if err != nil {
		return err
	}

	_, _, err = t.client.Statuses.Update(fmt.Sprintf("Song: %s\nAlbum: %s", song.Name, song.Album), &twitter.StatusUpdateParams{
		InReplyToStatusID: tweet.ID,
	})
	if err != nil {
		return err
	}

	return nil
}
