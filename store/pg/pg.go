package pg

import (
	"context"
	"errors"
	"fmt"
	"strings"
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

		log.Debug().Msg("Creating DB if not exists..")

		err = createDBIfNotExists(db)
		if err != nil {
			panic(err)
		}

		url := getUrlToDB()
		dbCreated, err := gorm.Open(postgres.Open(url), &gorm.Config{})
		if err != nil {
			log.Info().Msg("Connection to db failed")
			panic(err)
		}

		pgInstance = &DB{dbCreated}
	})
	return pgInstance, nil
}

func createDBIfNotExists(db *gorm.DB) error {
	var exists bool

	err := db.Raw("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'songs')").Scan(&exists).Error
	if err != nil {
		return err
	}

	if !exists {
		err := db.Exec("CREATE DATABASE songs").Error
		if err != nil {
			return err
		}
	}

	return nil
}

func getUrlToDB() string {
	conf := config.Get()
	return strings.ReplaceAll(conf.PgURL, "localhost/?", fmt.Sprintf("localhost/%s?", "songs"))
}
