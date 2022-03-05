package domain

import "math/rand"

type Song struct {
	ID     int      `json:"id" firestore:"id"`
	Name   string   `json:"name"  firestore:"name"`
	Album  string   `json:"album"  firestore:"album"`
	Lyrics []string `json:"lyrics"  firestore:"lyrics"`
}

func (s *Song) GetRandomLyric() string {
	return s.Lyrics[rand.Intn(len(s.Lyrics))]

}
