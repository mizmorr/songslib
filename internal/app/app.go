package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mizmorr/songslib/internal/controller"
	"github.com/mizmorr/songslib/internal/router"
	"github.com/mizmorr/songslib/pkg/logger"
	"github.com/mizmorr/songslib/pkg/server"
	"github.com/mizmorr/songslib/service"
	"github.com/mizmorr/songslib/store"
)

func Run() error {
	var (
		log       = logger.Get()
		ctx       = context.WithValue(context.Background(), "logger", log)
		interrupt = make(chan os.Signal, 1)
	)

	log.Debug().Msg("[app.Run] - store initialization...")
	store, err := store.New(ctx)
	if err != nil {
		return err
	}

	log.Debug().Msg("[app.Run] - service for songs initialization...")
	ws, err := service.NewSongWebService(ctx, store)
	if err != nil {
		return err
	}

	log.Debug().Msg("[app.Run] - controller initialization...")
	controller := controller.NewSongController(ctx, ws)

	log.Debug().Msg("[app.Run] - starting server...")
	handler := gin.New()
	router.NewRouter(handler, controller)

	httpServer := server.New(handler)

	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	log.Info().Msg("Server is running ...")

	select {
	case s := <-interrupt:
		log.Info().Msg("[app.Run] - signal " + s.String())
		// sort of graceful shutdown
		time.Sleep(500 * time.Millisecond)
	case err = <-httpServer.Notify():
		log.Error().Err(fmt.Errorf("[app.Run] - httpServer.Notify " + err.Error()))
	}

	return httpServer.Shutdown()
}
