package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	firebase "firebase.google.com/go"
	"github.com/jensmcatanho/avantasia-txt/client"
	"github.com/jensmcatanho/avantasia-txt/models"
	"github.com/labstack/echo/v4"
)

var (
	twitterClient *client.TwitterClient
	totalSongs    int
)

func main() {
	rand.Seed(time.Now().UnixNano())
	twitterClient = client.NewTwitterClient(client.Credentials{
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	})

	var err error
	totalSongs, err = strconv.Atoi(os.Getenv("TOTAL_SONGS"))
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.GET("/tweet", TweetHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func TweetHandler(echoContext echo.Context) error {
	song, err := getRandomSong(echoContext.Request().Context())
	if err != nil {
		log.Printf("Error: %+v", err)
		return echoContext.NoContent(http.StatusInternalServerError)
	}

	log.Printf("Song: %s\n", song.Name)
	err = twitterClient.Tweet(song)
	if err != nil {
		log.Printf("Error: %+v", err)
		return echoContext.NoContent(http.StatusInternalServerError)
	}

	return echoContext.NoContent(http.StatusOK)
}

func getRandomSong(ctx context.Context) (*models.Song, error) {
	conf := &firebase.Config{
		DatabaseURL: os.Getenv("FIREBASE_URL"),
	}

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	iter := client.Collection("songs").Where("id", "==", fmt.Sprint(rand.Intn(totalSongs)+1)).Limit(1).Documents(ctx)
	document, err := iter.Next()
	if err != nil {
		return nil, err
	}

	var song models.Song
	err = document.DataTo(&song)
	if err != nil {
		return nil, err
	}

	return &song, nil
}
