package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mizmorr/songslib/internal/controller"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(handler *gin.Engine, c *controller.SongController) {
	handler.Use(gin.Recovery())
	handler.Use(gin.Logger())
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))
	v1 := handler.Group("/v1")
	{
		songRoutes := v1.Group("/song")
		songRoutes.POST("", c.Create)
		songRoutes.DELETE("", c.Delete)
		songRoutes.PUT("", c.Update)
		songRoutes.GET("/verses", c.GetVersesOfSong)
		songRoutes.GET("/pages", c.GetAllFiltredPaginated)
	}
	// handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
