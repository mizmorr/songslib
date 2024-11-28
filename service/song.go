package service

import (
	"context"

	"github.com/mizmorr/songslib/internal/model"
	"github.com/mizmorr/songslib/store"
	"github.com/pkg/errors"
)

type SongWebService struct {
	ctx   context.Context
	store *store.Store
}

func NewSongWebService(ctx context.Context, store *store.Store) (*SongWebService, error) {
	if store == nil {
		return nil, errors.New("store cannot be nil")
	}
	return &SongWebService{
		ctx:   ctx,
		store: store,
	}, nil
}

func (ws *SongWebService) Create(ctx context.Context, song *model.SongRequestCreate) (id uint, err error) {
	songDB := song.ToDB()
	id, err = ws.store.Song.Create(ctx, songDB)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (ws *SongWebService) Delete(ctx context.Context, song *model.SongRequestDelete) error {
	songDB := song.ToDB()
	err := ws.store.Song.Delete(ctx, songDB)
	if err != nil {
		return err
	}
	return nil
}

func (ws *SongWebService) Update(ctx context.Context, song *model.SongRequestUpdate) error {
	songDB := song.ToDB()
	err := ws.store.Song.Update(ctx, songDB)
	if err != nil {
		return err
	}
	return nil
}

func (ws *SongWebService) GetVersesOfSong(ctx context.Context, song *model.SongRequestGet, pageOpts *model.Page) (*model.Verse, error) {
	verse, err := ws.store.Song.GetVerses(ctx, song.ToDB(), pageOpts)
	if err != nil {
		return nil, err
	}
	return verse, nil
}

func (ws *SongWebService) GetAllFiltredPaginated(ctx context.Context, song *model.SongRequestGet, pageOpts *model.Page) (int64, []*model.Song, error) {
	totalSongs, songs, err := ws.store.Song.GetAllFiltredPaginated(ctx, song.ToDB(), pageOpts)
	if err != nil {
		return totalSongs, nil, err
	}
	return totalSongs, songs, nil
}
