package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	songJson, err := os.Open("albums/the_wicked_symphony/1_the_wicked_symphony.json")
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}

	decoder := json.NewDecoder(songJson)

	var song Song
	err = decoder.Decode(&song)
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}

	fmt.Printf("%s\n", song.getLyric())
}

type Song struct {
	Name   string   `json:"name"`
	Album  string   `json:"album"`
	Lyrics []string `json:"lyrics"`
}

func (s *Song) getLyric() string {
	return s.Lyrics[rand.Intn(len(s.Lyrics))]
}
