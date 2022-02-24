package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jensmcatanho/avantasia-txt/models"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	songJson, err := os.Open("albums/the_wicked_symphony/1_the_wicked_symphony.json")
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}

	decoder := json.NewDecoder(songJson)

	var song models.Song
	err = decoder.Decode(&song)
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}

	fmt.Printf("%s\n", song.GetLyric())
}
