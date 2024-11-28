package controller

import (
	"context"
	"testing"

	"github.com/mizmorr/songslib/service"
	"github.com/mizmorr/songslib/store"
)

func TestNewSongController(t *testing.T) {
	ctx := context.Background()
	store, err := store.New(ctx)
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to connect to the database: %v", err)
	}

	sws, err := service.NewSongWebService(ctx, store)
	if err != nil {
		t.Errorf("[TestGetVerses] Expected a non-nil SongWebService, got %v+ ", err)
	}
	controller := NewSongController(ctx, sws)
	if controller == nil {
		t.Error("[TestNewSongController] Expected a non-nil controller, got nil")
	}
}
