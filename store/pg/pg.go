package pg

import (
	"errors"
	"sync"
	"time"

	"github.com/mizmorr/songslib/config"
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

func Dial() (*DB, error) {
	conf := config.Get()

	if conf.PgURL == "" {
		return nil, errors.New("PG_URL is not set")
	}

	// log:= logger.Get

	once.Do(func() {
		var (
			db  *gorm.DB
			err error
		)

		for conf.PgConnAttempts > 0 {
			db, err = gorm.Open(postgres.Open(conf.PgURL), &gorm.Config{})
			if err == nil {
				break
			}

			conf.PgConnAttempts--
			// logInfo
			time.Sleep(conf.PgTimeout)
		}
		if err != nil {
			// logError
			panic("Connection to PostgreSQL failed")
		}
		pgInstance = &DB{db}
	})

	return pgInstance, nil
}
