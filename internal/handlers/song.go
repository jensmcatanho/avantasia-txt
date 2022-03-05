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

type tweetLyricsQueryParams struct {
	Song    string `query:"song"`
	LyricID string `query:"lyricId"`
}

func (sh *songHandler) tweetLyrics(echoContext echo.Context) error {
	logger := sh.context.Logger()

	var queryParams tweetLyricsQueryParams
	if err := (&echo.DefaultBinder{}).BindQueryParams(echoContext, &queryParams); err != nil {
		logger.Error(err.Error())
	}

	song, err := sh.getSong(echoContext, queryParams)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	err = sh.tweetLyric(echoContext, song, queryParams)
	if err != nil {
		return err
	}

	logger.Info("Song tweeted successfully", zap.String("Song", song.Name))
	return echoContext.NoContent(http.StatusOK)
}

func (sh *songHandler) getSong(echoContext echo.Context, queryParams tweetLyricsQueryParams) (*domain.Song, error) {
	requestContext := echoContext.Request().Context()

	if len(queryParams.Song) == 0 {
		song, err := sh.songService.GetRandomSong(requestContext)
		if err != nil {
			return nil, echoContext.NoContent(http.StatusInternalServerError)
		}

		return song, nil
	}

	songName, err := url.QueryUnescape(queryParams.Song)
	if err != nil {
		return nil, echoContext.NoContent(http.StatusBadRequest)
	}

	song, err := sh.songService.GetSongByName(requestContext, songName)
	if err != nil {
		return nil, echoContext.NoContent(http.StatusInternalServerError)
	}

	return song, nil
}

func (sh *songHandler) tweetLyric(echoContext echo.Context, song *domain.Song, queryParams tweetLyricsQueryParams) error {
	if len(queryParams.LyricID) == 0 {
		err := sh.twitterService.TweetRandomLyric(song)
		if err != nil {
			return echoContext.NoContent(http.StatusInternalServerError)
		}

		return nil
	}

	err := sh.twitterService.TweetLyric(song, queryParams.LyricID)
	if err != nil {
		return echoContext.NoContent(http.StatusInternalServerError)
	}

	return nil
}
