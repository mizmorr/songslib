package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/mizmorr/songslib/internal/model"
	"github.com/mizmorr/songslib/store"
	"gorm.io/gorm"
)

func TestNewSongWebService(t *testing.T) {
	ctx := context.Background()

	_, err := NewSongWebService(ctx, nil)
	if err == nil {
		t.Error("[TestNewSongWebService] Expected a nil SongWebService")
	}
}

func TestCreate(t *testing.T) {
	ctx := context.Background()

	store, err := store.New(ctx)
	if err != nil {
		t.Errorf("[TestCreate] Expected a non-nil store, got %v+ ", err)
	}

	sws, err := NewSongWebService(ctx, store)
	if err != nil {
		t.Errorf("[TestCreate] Expected a non-nil SongWebService, got %v+ ", err)
	}
	songRequest := &model.SongRequestCreate{
		Name: "Test Song",
		Band: "Test Band",
	}

	id, err := sws.Create(ctx, songRequest)
	if err != nil {
		t.Errorf("[TestCreate] Failed to create song: %v", err)
	}

	if id == 0 {
		t.Error("[TestCreate] Failed to retrieve song ID")
	}
	songToDelete := &model.Song{
		Model: gorm.Model{
			ID: id,
		},
		Name: "Test Song",
		Band: "Test Band",
	}
	err = store.Pg.Unscoped().Delete(songToDelete).Error
	if err != nil {
		t.Errorf("[TestDelete] Failed to delete song: %v", err)
	}
}

func TestDelete(t *testing.T) {
	ctx := context.Background()

	store, err := store.New(ctx)
	if err != nil {
		t.Errorf("[TestDelete] Expected a non-nil store, got %v+ ", err)
	}
	songToDelete := &model.Song{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "Test Song",
	}
	_, err = store.Song.Create(ctx, songToDelete)
	if err != nil {
		t.Errorf("[TestDelete] Failed to create song: %v", err)
	}

	sws, err := NewSongWebService(ctx, store)
	if err != nil {
		t.Errorf("[TestDelete] Expected a non-nil SongWebService, got %v+ ", err)
	}
	songRequest := &model.SongRequestDelete{
		ID: 1,
	}

	err = sws.Delete(ctx, songRequest)
	if err != nil {
		t.Errorf("[TestDelete] Failed to delete song: %v", err)
	}

	err = store.Pg.Unscoped().Delete(&songToDelete).Error
	if err != nil {
		t.Errorf("[TestDelete] Failed to delete song: %v", err)
	}
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()

	store, err := store.New(ctx)
	if err != nil {
		t.Errorf("[TestUpdate] Expected a non-nil store, got %v+ ", err)
	}
	songToUpdate := &model.Song{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "Test Song",
	}
	_, err = store.Song.Create(ctx, songToUpdate)
	if err != nil {
		t.Errorf("[TestUpdate] Failed to create song: %v", err)
	}

	sws, err := NewSongWebService(ctx, store)
	if err != nil {
		t.Errorf("[TestUpdate] Expected a non-nil SongWebService, got %v+ ", err)
	}
	songRequest := &model.SongRequestUpdate{
		ID:   1,
		Name: "Updated Test Song",
		Band: "Updated Test Band",
	}

	err = sws.Update(ctx, songRequest)
	if err != nil {
		t.Errorf("[TestUpdate] Failed to update song: %v", err)
	}

	err = store.Pg.Unscoped().Delete(&songToUpdate).Error
	if err != nil {
		t.Errorf("[TestDelete] Failed to delete song: %v", err)
	}
}

func TestGetVerses(t *testing.T) {
	songGet := &model.SongRequestGet{
		Name:   "Yesterday",
		Band:   "The Beatles",
		Lyrics: "All my troubles seemed so far away,\nNow it looks as though they’re here to stay.\nOh, I believe in yesterday.",
	}
	pageOpts := &model.Page{
		Number: 1,
		Size:   1,
	}
	expectedVerse := &model.Verse{
		Song:   "Yesterday",
		Band:   "The Beatles",
		Number: 1,
		Lines:  []string{"All my troubles seemed so far away,"},
	}

	ctx := context.Background()

	store, err := store.New(ctx)
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to connect to the database: %v", err)
	}

	id, err := store.Song.Create(ctx, songGet.ToDB())

	sws, err := NewSongWebService(ctx, store)
	if err != nil {
		t.Errorf("[TestGetVerses] Expected a non-nil SongWebService, got %v+ ", err)
	}
	verse, err := sws.GetVersesOfSong(ctx, songGet, pageOpts)
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to get verses: %v", err)
	}

	if !reflect.DeepEqual(verse, expectedVerse) {
		t.Errorf("[TestGetVerses] Expected %+v, got %+v", expectedVerse, verse)
	}
	songGet.ID = id
	err = store.Pg.Unscoped().Delete(songGet.ToDB()).Error
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to delete song: %v", err)
	}
}

func TestGetFiltredPaginated(t *testing.T) {
	var (
		id      uint = 1
		songGet      = &model.SongRequestGet{
			ID:     id,
			Name:   "Yesterday",
			Band:   "The Beatles",
			Lyrics: "All my troubles seemed so far away,\nNow it looks as though they’re here to stay.\nOh, I believe in yesterday.",
		}
		songForFiltrPattern = &model.SongRequestGet{
			Name: "Yes",
			Band: "The Be",
		}
		pageOpts = &model.Page{
			Number: 1,
			Size:   1,
		}
		expectedReturnedCount int64 = 1
		ctx                         = context.Background()
	)
	store, err := store.New(ctx)
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to connect to the database: %v", err)
	}

	_, err = store.Song.Create(ctx, songGet.ToDB())

	sws, err := NewSongWebService(ctx, store)
	if err != nil {
		t.Errorf("[TestGetVerses] Expected a non-nil SongWebService, got %v+ ", err)
	}

	returnedCount, selectedSongs, err := sws.GetAllFiltredPaginated(ctx, songForFiltrPattern, pageOpts)
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to get verses: %v", err)
	}

	if returnedCount != expectedReturnedCount {
		t.Errorf("[TestGetVerses] Expected %d returned songs, got %d", expectedReturnedCount, returnedCount)
	}

	if selectedSongs[0].ID != id {
		t.Errorf("[TestGetVerses] Expected id %+v, got %+v", id, selectedSongs[0].ID)
	}

	songGet.ID = id
	err = store.Pg.Unscoped().Delete(songGet.ToDB()).Error
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to delete song: %v", err)
	}
}
