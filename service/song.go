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

/*
Create(ctx context.Context, song *model.Song) (id uint, err error)
	Delete(ctx context.Context, song *model.Song) error
	GetAllFiltredPaginated(ctx context.Context, song *model.Song, pageOpts *model.Page) (int64, []*model.Song, error)
	Update(ctx context.Context, song *model.Song) error
	GetVerses(ctx context.Context, song *model.Song, pageOpt model.Page) (*model.Verse, error)
*/

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
		return 0, errors.Wrap(err, "svc.Song.Create")
	}
	return id, nil
}

func (ws *SongWebService) Delete(ctx context.Context, song *model.SongRequestDelete) error {
	songDB := song.ToDB()
	err := ws.store.Song.Delete(ctx, songDB)
	if err != nil {
		return errors.Wrap(err, "svc.Song.Delete")
	}
	return nil
}

func (ws *SongWebService) Update(ctx context.Context, song *model.SongRequestUpdate) error {
	songDB := song.ToDB()
	err := ws.store.Song.Update(ctx, songDB)
	if err != nil {
		return errors.Wrap(err, "svc.Song.Update")
	}
	return nil
}

func (ws *SongWebService) GetVersesOfSong(ctx context.Context, song *model.Song, pageOpts *model.Page) (*model.Verse, error) {
	verse, err := ws.store.Song.GetVerses(ctx, song, pageOpts)
	if err != nil {
		return nil, errors.Wrap(err, "svc.Song.GetVerseOfSong")
	}
	return verse, nil
}

func (ws *SongWebService) GetAllFiltredPaginated(ctx context.Context, song *model.Song, pageOpts *model.Page) (int64, []*model.Song, error) {
	totalSongs, songs, err := ws.store.Song.GetAllFiltredPaginated(ctx, song, pageOpts)
	if err != nil {
		return totalSongs, nil, errors.Wrap(err, "svc.Song.GetAllFiltredPaginated")
	}
	return totalSongs, songs, nil
}
