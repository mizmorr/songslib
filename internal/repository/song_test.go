package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/mizmorr/songslib/internal/model"
	"github.com/mizmorr/songslib/store/pg"
	"gorm.io/gorm"
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

func TestUpdate(t *testing.T) {
	var (
		id          uint = 1
		song             = model.Song{Model: gorm.Model{ID: id}, Band: "The Beatles", Name: "Help!", Lyrics: "Help me, please, help me now\nWith a loving heart\nI've been broken, but I'm trying to find my way\nBack to the light"}
		updatedSong      = model.Song{Model: gorm.Model{ID: id}, Band: "The Beatles", Name: "Do not help", Lyrics: "Do not help me!"}
		ctx              = context.Background()
	)
	db, err := pg.Dial()
	if err != nil {
		t.Errorf("[TestUpdate] Failed to connect to the database: %s", err)
	}
	songRepo := NewSongRepository(db)

	returningID, err := songRepo.Create(ctx, song)
	if err != nil {
		t.Errorf("[TestUpdate] Failed to create song: %v", err)
	}

	if returningID != id {
		t.Errorf("[TestUpdate] Expected returning ID to be %d, got %d", id, returningID)
	}

	err = songRepo.Update(ctx, updatedSong)
	if err != nil {
		t.Errorf("[TestUpdate] Failed to update song: %v", err)
	}

	var afterUpdatingSong model.Song

	result := db.First(&afterUpdatingSong)

	if result.Error != nil {
		t.Errorf("[TestUpdate] Failed to get updated song: %v", result.Error)
	}
	if afterUpdatingSong.Band != updatedSong.Band || afterUpdatingSong.Lyrics != updatedSong.Lyrics || afterUpdatingSong.Name != updatedSong.Name {
		t.Errorf("[TestUpdate] Results don't match after update: %+v, expected %+v", afterUpdatingSong, updatedSong)
	}
	err = db.Unscoped().Delete(&song).Error
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to delete song: %v", err)
	}
}
