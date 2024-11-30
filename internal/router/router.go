package router

import (
	"song-library/internal/handler"
	"song-library/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, songService *service.SongService) {

	api := r.Group("/api/v1/songs")
	{
		api.GET("", handler.GetSongs(songService))
		api.GET("/:id", handler.GetSongByID(songService))
		api.POST("", handler.AddSong(songService))
		api.PUT("/:id", handler.UpdateSong(songService))
		api.DELETE("/:id", handler.DeleteSong(songService))
	}
}
