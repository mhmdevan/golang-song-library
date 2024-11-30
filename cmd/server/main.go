package main

import (
	"song-library/config"
	"song-library/internal/db"
	"song-library/internal/repository"
	"song-library/internal/router"
	"song-library/internal/service"
	"song-library/pkg/logger"

	_ "song-library/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Song Library API
// @version 1.0
// @description This is the API documentation for the Song Library service.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Initialize logger
	logger.Init()

	// Load configuration
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		logger.Error("Failed to load configuration", logger.Fields{"error": err.Error()})
		return
	}

	// Initialize database connection
	dbConn, err := db.Connect(cfg)
	if err != nil {
		logger.Error("Failed to connect to database", logger.Fields{"error": err.Error()})
		return
	}
	defer func() {
		sqlDB, _ := dbConn.DB()
		sqlDB.Close()
		logger.Info("Database connection closed", nil)
	}()

	// Initialize repositories and services
	songRepository := repository.NewSongRepository(dbConn)
	songService := service.NewSongService(songRepository)

	// Initialize Gin engine
	r := gin.Default()

	// Setup Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes with services
	router.SetupRoutes(r, songService)

	// Start the server
	logger.Info("Server is starting", logger.Fields{"port": cfg.ServerPort})
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		logger.Error("Failed to start server", logger.Fields{"error": err.Error()})
	}
}
