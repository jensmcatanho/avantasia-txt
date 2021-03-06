package repositories

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/jensmcatanho/avantasia-txt/internal/core/domain"
)

type songRepository struct {
	client *firestore.Client
}

func NewSongRepository() (*songRepository, error) {
	ctx := context.Background()

	firebaseConfig := &firebase.Config{
		DatabaseURL: os.Getenv("FIREBASE_URL"),
	}

	app, err := firebase.NewApp(ctx, firebaseConfig)
	if err != nil {
		return nil, err
	}

	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return &songRepository{
		client: firestoreClient,
	}, nil
}

func (sr *songRepository) GetSongByID(ctx context.Context, id int) (*domain.Song, string, error) {
	return sr.getSong(ctx, "id", id)
}

func (sr *songRepository) GetSongByName(ctx context.Context, name string) (*domain.Song, string, error) {
	return sr.getSong(ctx, "name", name)
}

func (sr *songRepository) getSong(ctx context.Context, searchField string, value interface{}) (*domain.Song, string, error) {
	iter := sr.client.Collection("songs").Where(searchField, "==", value).Documents(ctx)
	document, err := iter.Next()
	if err != nil {
		return nil, "", err
	}

	var song domain.Song
	err = document.DataTo(&song)
	if err != nil {
		return nil, "", err
	}

	return &song, document.Ref.ID, nil
}

func (sr *songRepository) UpdateSong(ctx context.Context, referenceID string, song *domain.Song) error {
	_, err := sr.client.Collection("songs").Doc(referenceID).Update(ctx, []firestore.Update{{
		Path:  "lyrics",
		Value: song.Lyrics,
	}})

	return err
}
