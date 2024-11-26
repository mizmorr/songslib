package main

import (
	"context"
	"log"

	"github.com/mizmorr/songslib/internal/model"
	"github.com/mizmorr/songslib/internal/repository"
	"github.com/mizmorr/songslib/store/pg"
)

func main() {
	song := model.Song{
		Band: "The Beatles",
		Name: "Yesterday",
	}

	ctx := context.Background()
	db, err := pg.Dial()
	if err != nil {
		log.Fatal(err)
	}
	songRepository := repository.NewSongRepository(db)

	id, err := songRepository.Create(ctx, song)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Song created with ID: %d\n", id)

	err = songRepository.Delete(ctx, song)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Song deleted successfully")
}
