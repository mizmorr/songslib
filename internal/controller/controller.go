package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mizmorr/songslib/internal/model"
	e "github.com/mizmorr/songslib/pkg/errors"
	"github.com/mizmorr/songslib/service"
)

type WebService interface {
	Create(ctx context.Context, song *model.SongRequestCreate) (id uint, err error)
	Delete(ctx context.Context, song *model.SongRequestDelete) error
	Update(ctx context.Context, song *model.SongRequestUpdate) error
	GetVersesOfSong(ctx context.Context, song *model.SongRequestGet, pageOpts *model.Page) (*model.Verse, error)
	GetAllFiltredPaginated(ctx context.Context, song *model.SongRequestGet, pageOpts *model.Page) (int64, []*model.Song, error)
}

var _ WebService = (*service.SongWebService)(nil)

type SongController struct {
	ctx context.Context
	ws  WebService
}

func NewSongController(ctx context.Context, ws WebService) *SongController {
	return &SongController{
		ctx: ctx,
		ws:  ws,
	}
}

func (sc *SongController) Create(g *gin.Context) {
	song := &model.SongRequestCreate{}

	if err := g.BindJSON(song); err != nil {
		return
	}

	createdID, err := sc.ws.Create(sc.ctx, song)
	if err != nil {
		respondWithError(g, e.ErrInternalServer, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": createdID})
}

func (sc *SongController) Delete(g *gin.Context) {
	song := &model.SongRequestDelete{}

	if err := g.BindJSON(song); err != nil {
		return
	}

	err := sc.ws.Delete(sc.ctx, song)
	if err != nil {
		respondWithError(g, e.ErrBadRequest, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"status": "successfully deleted"})
}

func (sc *SongController) Update(g *gin.Context) {
	song := &model.SongRequestUpdate{}

	if err := g.BindJSON(song); err != nil {
		return
	}

	err := sc.ws.Update(sc.ctx, song)
	if err != nil {
		respondWithError(g, e.ErrBadRequest, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"status": "successfully updated"})
}

func (sc *SongController) GetVersesOfSong(g *gin.Context) {
	var (
		song     = &model.SongRequestGet{}
		pageOpts = &model.Page{}
	)

	if err := g.ShouldBindJSON(song); err != nil {
		respondWithError(g, e.ErrBadRequest, err.Error())
		return
	}

	if err := g.ShouldBindQuery(pageOpts); err != nil {
		respondWithError(g, e.ErrBadRequest, err.Error())
		return
	}

	verse, err := sc.ws.GetVersesOfSong(sc.ctx, song, pageOpts)
	if err != nil {
		respondWithError(g, e.ErrInternalServer, err.Error())
		return
	}
	g.JSON(http.StatusOK, verse)
}

func (sc *SongController) GetAllFiltredPaginated(g *gin.Context) {
	var (
		song     = &model.SongRequestGet{}
		pageOpts = &model.Page{}
	)

	if err := g.ShouldBindJSON(song); err != nil {
		respondWithError(g, e.ErrBadRequest, err.Error())
		return
	}

	if err := g.ShouldBindQuery(pageOpts); err != nil {
		respondWithError(g, e.ErrBadRequest, err.Error())
		return
	}
	totalSongs, songs, err := sc.ws.GetAllFiltredPaginated(sc.ctx, song, pageOpts)
	if err != nil {
		respondWithError(g, e.ErrInternalServer, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"total_songs": totalSongs, "songs": songs})
}

func respondWithError(g *gin.Context, err e.AppError, details string) {
	err.Details = details
	g.AbortWithStatusJSON(err.StatusCode, err)
}
