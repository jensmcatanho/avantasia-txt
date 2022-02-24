package models

import "math/rand"

type Song struct {
	Name   string   `json:"name"`
	Album  string   `json:"album"`
	Lyrics []string `json:"lyrics"`
}

func (s *Song) GetLyric() string {
	return s.Lyrics[rand.Intn(len(s.Lyrics))]
}
