package handlers

import (
	"net/http"

	"github.com/jensmcatanho/avantasia-txt/internal/core/domain"
	"github.com/jensmcatanho/avantasia-txt/internal/core/ports"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type songHandler struct {
	context        *domain.AppContext
	echo           *echo.Echo
	songService    ports.SongService
	twitterService ports.TwitterService
}

func NewSongsHandler(ctx *domain.AppContext, e *echo.Echo, songService ports.SongService, twitterService ports.TwitterService) *songHandler {
	return &songHandler{
		context:        ctx,
		echo:           e,
		songService:    songService,
		twitterService: twitterService,
	}
}

func (sh *songHandler) SetupEndpoints() {
	sh.echo.GET("/tweet", sh.tweetLyrics)
}

func (sh *songHandler) tweetLyrics(echoContext echo.Context) error {
	logger := sh.context.Logger()
	song, err := sh.songService.GetRandomSong(echoContext.Request().Context())
	if err != nil {
		logger.Error(err.Error())
		return echoContext.NoContent(http.StatusInternalServerError)
	}

	err = sh.twitterService.Tweet(song)
	if err != nil {
		logger.Error(err.Error())
		return echoContext.NoContent(http.StatusInternalServerError)
	}

	logger.Info("Song tweeted successfully", zap.String("Song", song.Name))
	return echoContext.NoContent(http.StatusOK)
}
