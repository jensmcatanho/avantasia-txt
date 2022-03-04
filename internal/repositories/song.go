package repositories

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/jensmcatanho/avantasia-txt/internal/core/domain"
	"google.golang.org/api/option"
)

type songRepository struct {
	client *firestore.Client
}

func NewSongRepository() (*songRepository, error) {
	ctx := context.Background()

	firebaseConfig := &firebase.Config{
		DatabaseURL: os.Getenv("FIREBASE_URL"),
	}
	opt := option.WithCredentialsFile("key.json")

	app, err := firebase.NewApp(ctx, firebaseConfig, opt)
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

func (sr *songRepository) GetSongByID(ctx context.Context, id int) (*domain.Song, error) {
	iter := sr.client.Collection("songs").Where("id", "==", id).Documents(ctx)
	document, err := iter.Next()
	if err != nil {
		return nil, err
	}

	var song domain.Song
	err = document.DataTo(&song)
	if err != nil {
		return nil, err
	}

	return &song, nil
}
