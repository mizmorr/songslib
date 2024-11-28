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
	ctx := context.Background()
	db, err := pg.Dial(ctx)
	if err != nil {
		t.Errorf("[TestNewSongRepository] Failed to connect to the database: %v", err)
	}
	songRepo := NewSongRepository(db)
	if songRepo == nil {
		t.Errorf("[TestNewSongRepository] Expected a non-nil SongRepository")
	}
}

func TestCreate(t *testing.T) {
	var (
		song            = &model.Song{Model: gorm.Model{ID: 1}, Band: "The Beatles", Name: "Help!", Lyrics: "Help me, please, help me now\nWith a loving heart\nI've been broken, but I'm trying to find my way\nBack to the light"}
		ctx             = context.Background()
		expectedID uint = 1
	)
	db, err := pg.Dial(ctx)
	if err != nil {
		t.Errorf("[TestCreate] Failed to connect to the database: %v", err)
	}
	songRepo := NewSongRepository(db)
	id, err := songRepo.Create(ctx, song)
	if err != nil {
		t.Errorf("[TestCreate] Failed to create song: %v", err)
	}
	if id != expectedID {
		t.Errorf("[TestCreate] Expected ID to be %d, got %d", expectedID, id)
	}
	err = db.Unscoped().Delete(song).Error
	if err != nil {
		t.Errorf("[TestCreate] Failed to delete song: %v", err)
	}
}

func TestGetVerses(t *testing.T) {
	var (
		song         = &model.Song{Band: "The Beatles", Name: "Help!", Lyrics: "Help me, please, help me now\nWith a loving heart\nI've been broken, but I'm trying to find my way\nBack to the light"}
		pageOpt      = &model.Page{Number: 2, Size: 1}
		ctx          = context.Background()
		targetResult = &model.Verse{Number: 2, Song: "Help!", Band: "The Beatles", Lines: []string{"With a loving heart"}}
	)
	db, err := pg.Dial(ctx)
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

	err = db.Unscoped().Delete(song).Error
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to delete song: %v", err)
	}
}

func TestGetVersesBadSize(t *testing.T) {
	var (
		song    = &model.Song{Band: "The Beatles", Name: "Help!", Lyrics: "Help me, please, help me now\nWith a loving heart\nI've been broken, but I'm trying to find my way\nBack to the light"}
		pageOpt = &model.Page{Number: 3, Size: 3}
		ctx     = context.Background()
	)
	db, err := pg.Dial(ctx)
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to connect to the database: %v", err)
	}
	songRepo := NewSongRepository(db)
	id, err := songRepo.Create(ctx, song)
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to create song: %v", err)
	}
	_, err = songRepo.GetVerses(ctx, song, pageOpt)
	if err == nil {
		t.Error("[TestGetVerses] Expected an error when size is greater than the total number of verses")
	}
	{
		song.ID = id
		err = db.Unscoped().Delete(song).Error
		if err != nil {
			t.Errorf("[TestGetVerses] Failed to delete song: %v", err)
		}
	}
}

func TestUpdate(t *testing.T) {
	var (
		id          uint = 1
		song             = &model.Song{Model: gorm.Model{ID: id}, Band: "The Beatles", Name: "Help!", Lyrics: "Help me, please, help me now\nWith a loving heart\nI've been broken, but I'm trying to find my way\nBack to the light"}
		updatedSong      = &model.Song{Model: gorm.Model{ID: id}, Band: "The Beatles", Name: "Do not help", Lyrics: "Do not help me!"}
		ctx              = context.Background()
	)
	db, err := pg.Dial(ctx)
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
	err = db.Unscoped().Delete(song).Error
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to delete song: %v", err)
	}
}

func TestGetAllFilteredPaginated(t *testing.T) {
	var (
		song                   = &model.Song{Band: "The Beatles", Name: "Help!", Lyrics: "Help me, please, help me now\nWith a loving heart\nI've been broken, but I'm trying to find my way\nBack to the light"}
		ctx                    = context.Background()
		songForFilter          = &model.Song{Band: "The Be", Name: "Help!"}
		targetCountSongs int64 = 1
	)
	db, err := pg.Dial(ctx)
	if err != nil {
		t.Errorf("[TestGetAllFilteredPaginated] Failed to connect to the database: %v", err)
	}
	songRepo := NewSongRepository(db)
	id, err := songRepo.Create(ctx, song)
	if err != nil {
		t.Errorf("[TestGetAllFilteredPaginated] Failed to create song: %v", err)
	}
	pageOpts := &model.Page{Number: 1, Size: 1}
	totalCount, selectedSongs, err := songRepo.GetAllFiltredPaginated(ctx, songForFilter, pageOpts)
	if err != nil {
		t.Errorf("[TestGetAllFilteredPaginated] Failed to get filtered songs: %v", err)
	}
	if totalCount != targetCountSongs {
		t.Errorf("[TestGetAllFilteredPaginated] Expected total count to be %d, got %d", targetCountSongs, totalCount)
	}

	if selectedSongs[0].ID != id {
		t.Errorf("[TestGetAllFilteredPaginated] Expected selected song to have ID %d, got %d", id, selectedSongs[0].ID)
	}

	err = db.Unscoped().Delete(song).Error
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to delete song: %v", err)
	}
}

func TestGetAllFilteredPaginatedNoEnoughSongs(t *testing.T) {
	var (
		song                     = &model.Song{Band: "The Beatles", Name: "Help!", Lyrics: "Help me, please, help me now\nWith a loving heart\nI've been broken, but I'm trying to find my way\nBack to the light"}
		ctx                      = context.Background()
		songForFilter            = &model.Song{Band: "The Beatles", Name: "Help!"}
		pageOpts                 = &model.Page{Number: 2, Size: 2}
		expectedCountSongs int64 = 1
	)
	db, err := pg.Dial(ctx)
	if err != nil {
		t.Errorf("[TestGetAllFilteredPaginatedNoEnoughSongs] Failed to connect to the database: %v", err)
	}
	songRepo := NewSongRepository(db)
	_, err = songRepo.Create(ctx, song)
	if err != nil {
		t.Errorf("[TestGetAllFilteredPaginatedNoEnoughSongs] Failed to create song: %v", err)
	}
	totalCount, _, err := songRepo.GetAllFiltredPaginated(ctx, songForFilter, pageOpts)

	if err == nil {
		t.Error("[TestGetAllFilteredPaginatedNoEnoughSongs] Expected an error when there are no enough songs")
	}
	if totalCount != expectedCountSongs {
		t.Errorf("[TestGetAllFilteredPaginatedNoEnoughSongs] Expected total count to be %d, got %d", expectedCountSongs, totalCount)
	}

	err = db.Unscoped().Delete(song).Error
	if err != nil {
		t.Errorf("[TestGetVerses] Failed to delete song: %v", err)
	}
}
