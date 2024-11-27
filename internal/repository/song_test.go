package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/mizmorr/songslib/internal/model"
	"github.com/mizmorr/songslib/store/pg"
)

func TestNewSongRepository(t *testing.T) {
	db, err := pg.Dial()
	if err != nil {
		t.Errorf("[TestNewSongRepository] Failed to connect to the database: %v", err)
	}
	songRepo := NewSongRepository(db)
	if songRepo == nil {
		t.Errorf("[TestNewSongRepository] Expected a non-nil SongRepository")
	}
}

func TestGetVerses(t *testing.T) {
	var (
		song         = model.Song{Band: "The Beatles", Name: "Help!", Lyrics: "Help me, please, help me now\nWith a loving heart\nI've been broken, but I'm trying to find my way\nBack to the light"}
		pageOpt      = model.Page{Number: 2, Size: 1}
		ctx          = context.Background()
		targetResult = &model.Verse{Number: 2, Song: "Help!", Band: "The Beatles", Lines: []string{"With a loving heart"}}
	)
	db, err := pg.Dial()
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to connect to the database: %v", err)
	}
	songRepo := NewSongRepository(db)
	id, err := songRepo.Create(ctx, song)
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to create song: %v", err)
	}
	verse, err := songRepo.GetVerses(ctx, song, pageOpt)
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to get verses: %v", err)
	}
	if !reflect.DeepEqual(verse, targetResult) {
		t.Errorf("[TestGetVerses] Expected %+v, got %+v", targetResult, verse)
	}
	song.ID = id

	err = db.Unscoped().Delete(&song).Error
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to delete song: %v", err)
	}
}
