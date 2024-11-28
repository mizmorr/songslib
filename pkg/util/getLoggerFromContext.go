package util

import (
	"context"

	"github.com/mizmorr/songslib/pkg/logger"
)

func GetLogger(ctx context.Context) *logger.Logger {
	log, ok := ctx.Value("logger").(*logger.Logger)
	if !ok {
		return logger.Get()
	}
	return log
}
