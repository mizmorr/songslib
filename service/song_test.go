package service

import (
	"context"
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
