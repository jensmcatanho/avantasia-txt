package client

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
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

func (t *TwitterClient) Tweet(tweet string) error {
	_, _, err := t.client.Statuses.Update(tweet, nil)
	return err
}
