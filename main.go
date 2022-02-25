package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/jensmcatanho/avantasia-txt/client"
	"github.com/jensmcatanho/avantasia-txt/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	twitterClient := client.NewTwitterClient(client.Credentials{
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	})

	song, err := getRandomSong()
	if err != nil {
		log.Printf("Error: %+v", err)
	}

	log.Printf("Song: %s\n", song.Name)
	err = twitterClient.Tweet(song.GetLyric())
	if err != nil {
		log.Printf("Error: %+v", err)

	}
}

func getRandomSong() (*models.Song, error) {
	album := "the_wicked_symphony"

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
