package repository

import (
	"context"

	"github.com/mizmorr/songslib/internal/model"
	"github.com/mizmorr/songslib/store/pg"
)

type SongRepository struct {
	db *pg.DB
}

func NewSongRepository(db *pg.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (r *SongRepository) Create(ctx context.Context, model model.Song) (id uint, err error) {
	result := r.db.Create(&model)

	if result.Error != nil {
		return 0, result.Error
	}

	return model.ID, nil
}

func (r *SongRepository) Delete(ctx context.Context, model model.Song) error {
	result := r.db.Where("band=?", model.Band).Where("name=?", model.Name).Delete(&model)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
