package main

import "github.com/mizmorr/songslib/internal/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
