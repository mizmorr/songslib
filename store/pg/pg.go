package pg

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/mizmorr/songslib/config"
	"github.com/mizmorr/songslib/pkg/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

var (
	pgInstance *DB
	once       sync.Once
)

func Dial(ctx context.Context) (*DB, error) {
	conf := config.Get()
	log := util.GetLogger(ctx)
	if conf.PgURL == "" {
		return nil, errors.New("PG_URL is not set")
	}

	once.Do(func() {
		var (
			db  *gorm.DB
			err error
		)

		for conf.PgConnAttempts > 0 {
			db, err = gorm.Open(postgres.Open(conf.PgURL), &gorm.Config{})
			if err == nil {
				log.Info().Msg("Connected to PostgreSQL")
				break
			}

			conf.PgConnAttempts--
			log.Debug().Msg(fmt.Sprintf("Postgres is trying to connect, attempts left: %d", conf.PgConnAttempts))

			time.Sleep(conf.PgTimeout)
		}
		if err != nil {
			panic("Connection to PostgreSQL failed")
		}
		pgInstance = &DB{db}
	})

	return pgInstance, nil
}
