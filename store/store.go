package store

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/mizmorr/songslib/internal/model"
	"github.com/mizmorr/songslib/internal/repository"
	"github.com/mizmorr/songslib/pkg/util"
	"github.com/mizmorr/songslib/store/migrations"
	"github.com/mizmorr/songslib/store/pg"
)

type SongRepo interface {
	Create(ctx context.Context, song *model.Song) (id uint, err error)
	Delete(ctx context.Context, song *model.Song) error
	GetAllFiltredPaginated(ctx context.Context, song *model.Song, pageOpts *model.Page) (int64, []*model.Song, error)
	Update(ctx context.Context, song *model.Song) error
	GetVerses(ctx context.Context, song *model.Song, pageOpt *model.Page) (*model.Verse, error)
}

var _ SongRepo = (*repository.SongRepository)(nil)

type Store struct {
	Pg   *pg.DB
	Song SongRepo
}

var store Store

func New(ctx context.Context) (*Store, error) {
	logger := util.GetLogger(ctx)

	logger.Debug().Msg("Initializing PostgreSQL store")
	pg, err := pg.Dial(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "pg.Store: failed to connect to the database")
	}

	logger.Debug().Msg("Running PostgreSQL migrations")
	if err := migrations.AutoMigrate(pg.DB); err != nil {
		return nil, errors.Wrap(err, "pg.Store: failed to run migrations")
	}

	if pg != nil {
		store.Pg = pg
		go store.keepAlive(ctx)
		store.Song = repository.NewSongRepository(pg)
	}
	logger.Info().Msg("PostgreSQL store initialized successfully")
	return &store, nil
}

const KeepALiveTimeout = 5

func (store *Store) keepAlive(ctx context.Context) {
	logger := util.GetLogger(ctx)
	for {
		time.Sleep(time.Second * KeepALiveTimeout)
		var (
			lost_connection bool
			err             error
		)

		if store.Pg == nil {
			lost_connection = true
		}
		if lost_connection {
			logger.Debug().Msg("[store.keepAlive] Lost connection, is trying to reconnect...")
			store.Pg, err = pg.Dial(ctx)
			if err != nil {
				logger.Err(err)
			} else {
				logger.Debug().Msg("[store.keepAlive] Connection established")
			}
		}

	}
}
