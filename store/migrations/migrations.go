package migrations

import (
	"context"

	"github.com/mizmorr/songslib/internal/model"
	"github.com/mizmorr/songslib/pkg/util"
	"gorm.io/gorm"
)

func AutoMigrate(ctx context.Context, db *gorm.DB) error {
	logger := util.GetLogger(ctx)

	logger.Debug().Msg("Running migrations..")

	err := db.AutoMigrate(&model.Song{})
	if err != nil {
		return err
	}

	for _, song := range returnSongs() {
		db.Create(&song)
	}
	logger.Info().Msg("Migrations completed successfully.")
	return nil
}
