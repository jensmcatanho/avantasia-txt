package handlers

import (
	"net/http"
	"net/url"

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
	requestContext := echoContext.Request().Context()
	songName := echoContext.QueryParam("song")

	var song *domain.Song
	var err error
	if songName == "" {
		song, err = sh.songService.GetRandomSong(requestContext)
		if err != nil {
			logger.Error(err.Error())
			return echoContext.NoContent(http.StatusInternalServerError)
		}
	} else {
		songNameDecoded, err := url.QueryUnescape(songName)
		if err != nil {
			logger.Error(err.Error())
			return echoContext.NoContent(http.StatusBadRequest)
		}

		song, err = sh.songService.GetSongByName(requestContext, songNameDecoded)
		if err != nil {
			logger.Error(err.Error())
			return echoContext.NoContent(http.StatusInternalServerError)
		}
	}

	err = sh.twitterService.Tweet(song)
	if err != nil {
		logger.Error(err.Error())
		return echoContext.NoContent(http.StatusInternalServerError)
	}

	logger.Info("Song tweeted successfully", zap.String("Song", song.Name))
	return echoContext.NoContent(http.StatusOK)
}
