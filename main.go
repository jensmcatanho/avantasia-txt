package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/jensmcatanho/avantasia-txt/client"
	"github.com/jensmcatanho/avantasia-txt/models"
	"github.com/labstack/echo/v4"
)

var twitterClient *client.TwitterClient

func main() {
	rand.Seed(time.Now().UnixNano())
	twitterClient = client.NewTwitterClient(client.Credentials{
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	})

	e := echo.New()
	e.GET("/tweet", TweetHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

func TweetHandler(context echo.Context) error {
	song, err := getRandomSong()
	if err != nil {
		log.Printf("Error: %+v", err)
		return context.NoContent(http.StatusInternalServerError)
	}

	log.Printf("Song: %s\n", song.Name)
	err = twitterClient.Tweet(song)
	if err != nil {
		log.Printf("Error: %+v", err)
		return context.NoContent(http.StatusInternalServerError)
	}

	return context.NoContent(http.StatusOK)
}

func getRandomSong() (*models.Song, error) {
	albumFolders, err := ioutil.ReadDir("albums")
	if err != nil {
		log.Fatal(err)
	}
	album := albumFolders[rand.Intn(len(albumFolders))].Name()

	songFiles, err := ioutil.ReadDir(fmt.Sprintf("albums/%s", album))
	if err != nil {
		log.Fatal(err)
	}

	songFile := songFiles[rand.Intn(len(songFiles))]
	songJson, err := os.Open(fmt.Sprintf("albums/%s/%s", album, songFile.Name()))
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(songJson)

	var song models.Song
	err = decoder.Decode(&song)
	if err != nil {
		return nil, err
	}

	return &song, nil
}
