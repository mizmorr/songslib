package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mizmorr/songslib/internal/model"
	"github.com/mizmorr/songslib/pkg/errors"
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

// Create godoc
// @Summary Create a new song
// @Description Create a new song record in the database with the provided details.
// @Tags songs
// @Accept  json
// @Produce  json
// @Param song body model.SongRequestCreate true "Song data to be created"
// @Success 200 {object} map[string]string "Successfully created song"
// @Failure 400 {object} errors.AppError "Invalid input or request"
// @Failure 500 {object} errors.AppError "Internal server error"
// @Router /song [post]
func (sc *SongController) Create(g *gin.Context) {
	song := &model.SongRequestCreate{}

	if err := g.BindJSON(song); err != nil {
		return
	}

	createdID, err := sc.ws.Create(sc.ctx, song)
	if err != nil {
		respondWithError(g, errors.ErrInternalServer, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"id": createdID})
}

// Delete godoc
// @Summary Delete a song
// @Description Delete a song by its ID
// @Tags songs
// @Accept json
// @Produce json
// @Param song body model.SongRequestDelete true "Song to delete"
// @Success 200 {object} map[string]string "Successfully deleted"
// @Failure 400 {object} errors.AppError "Bad request, invalid data"
// @Router /song [delete]
func (sc *SongController) Delete(g *gin.Context) {
	song := &model.SongRequestDelete{}

	if err := g.BindJSON(song); err != nil {
		return
	}

	err := sc.ws.Delete(sc.ctx, song)
	if err != nil {
		respondWithError(g, errors.ErrBadRequest, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"status": "successfully deleted"})
}

// Update godoc
// @Summary Update song information
// @Description Update an existing song record with the provided details.
// @Tags songs
// @Accept  json
// @Produce  json
// @Param song body model.SongRequestUpdate true "Song data to be updated"
// @Success 200 {object} map[string]string "Successfully updated"
// @Failure 400 {object} errors.AppError "Invalid input or request"
// @Failure 500 {object} errors.AppError "Internal server error"
// @Router /song [put]
func (sc *SongController) Update(g *gin.Context) {
	song := &model.SongRequestUpdate{}

	if err := g.BindJSON(song); err != nil {
		return
	}

	err := sc.ws.Update(sc.ctx, song)
	if err != nil {
		respondWithError(g, errors.ErrBadRequest, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"status": "successfully updated"})
}

// GetVersesOfSong godoc
// @Summary Get verses of a song
// @Description Retrieve all verses of a song based on the song details and pagination options provided.
// @Tags songs
// @Accept  json
// @Produce  json
// @Param song query model.SongRequestGet true "Song details for which verses need to be fetched"
// @Param pageOpts query model.Page false "Pagination options"
// @Success 200 {array} model.Verse "List of song verses"
// @Failure 400 {object} errors.AppError "Invalid input or request"
// @Failure 500 {object} errors.AppError "Internal server error"
// @Router /song/verses [get]
func (sc *SongController) GetVersesOfSong(g *gin.Context) {
	var (
		song     = &model.SongRequestGet{}
		pageOpts = &model.Page{}
	)

	if err := g.ShouldBindQuery(song); err != nil {
		respondWithError(g, errors.ErrBadRequest, err.Error())
		return
	}
	if err := g.ShouldBindQuery(pageOpts); err != nil {
		respondWithError(g, errors.ErrBadRequest, err.Error())
		return
	}

	verse, err := sc.ws.GetVersesOfSong(sc.ctx, song, pageOpts)
	if err != nil {
		respondWithError(g, errors.ErrInternalServer, err.Error())
		return
	}

	g.JSON(http.StatusOK, verse)
}

// GetAllFiltredPaginated godoc
// @Summary Get filtered and paginated list of songs
// @Description Retrieve a list of songs based on filtering criteria and pagination options.
// @Tags songs
// @Accept  json
// @Produce  json
// @Param song query model.SongRequestGet true "Filtering criteria for the songs"
// @Param pageOpts query model.Page false "Pagination options"
// @Success 200 {object} map[string]string "Successfully retrieved filtered songs list"
// @Failure 400 {object} errors.AppError "Invalid input or request"
// @Failure 500 {object} errors.AppError "Internal server error"
// @Router /song/pages [get]
func (sc *SongController) GetAllFiltredPaginated(g *gin.Context) {
	var (
		song     = &model.SongRequestGet{}
		pageOpts = &model.Page{}
	)

	if err := g.ShouldBindQuery(song); err != nil {
		respondWithError(g, errors.ErrBadRequest, err.Error())
		return
	}

	if err := g.ShouldBindQuery(pageOpts); err != nil {
		respondWithError(g, errors.ErrBadRequest, err.Error())
		return
	}

	totalSongs, songs, err := sc.ws.GetAllFiltredPaginated(sc.ctx, song, pageOpts)
	if err != nil {
		respondWithError(g, errors.ErrInternalServer, err.Error())
		return
	}

	g.JSON(http.StatusOK, gin.H{"total_songs": totalSongs, "songs": songs})
}

func respondWithError(g *gin.Context, err errors.AppError, details string) {
	err.Details = details
	g.AbortWithStatusJSON(err.StatusCode, err)
}
