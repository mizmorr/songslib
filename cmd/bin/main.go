package main

import (
	_ "github.com/mizmorr/songslib/docs"
	"github.com/mizmorr/songslib/internal/app"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is example of RESTful api

//	@host		localhost:8080
//	@BasePath	/v1

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
