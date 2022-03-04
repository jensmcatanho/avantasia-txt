package cmd

import (
	"math/rand"
	"time"

	"github.com/jensmcatanho/avantasia-txt/internal/core/domain"
	"github.com/jensmcatanho/avantasia-txt/internal/core/services"
	"github.com/jensmcatanho/avantasia-txt/internal/handlers"
	"github.com/jensmcatanho/avantasia-txt/internal/repositories"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type AvantasiaTxt struct {
	Context domain.AppContext
}

func NewAvantasiaTxt(logger *zap.Logger) *AvantasiaTxt {
	appContext := domain.NewAppContext(*logger)

	return &AvantasiaTxt{
		Context: *appContext,
	}
}

func (a *AvantasiaTxt) Run() error {
	rand.Seed(time.Now().UnixNano())

	logger := a.Context.Logger()
	logger.Info("Starting application")

	e := echo.New()

	songRepository, err := repositories.NewSongRepository()
	if err != nil {
		return err
	}
	songService := services.NewSongService(songRepository)

	twitterCredentials := domain.NewTwitterCredentials()
	twitterService := services.NewTwitterService(twitterCredentials)

	songHandler := handlers.NewSongsHandler(&a.Context, e, songService, twitterService)
	songHandler.SetupEndpoints()

	err = e.Start(":8080")
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}
