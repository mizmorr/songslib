package migrations

import (
	"github.com/mizmorr/songslib/internal/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Song{})
	if err != nil {
		return err
	}

	return nil
}
