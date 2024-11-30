package test

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"song-library/config"
	"song-library/internal/db"
	"song-library/internal/model"
	"song-library/internal/repository"
	"song-library/internal/router"
	"song-library/internal/service"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
)

func setupIntegrationTest() (*gin.Engine, error) {
	// Load configuration
	cfg, err := config.LoadConfig("../../.env")
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %v", err)
	}

	// Connect to the database
	dbConn, err := db.Connect(cfg)
	if err != nil {
		if strings.Contains(err.Error(), "does not exist") {
			err = createDatabase(cfg)
			if err != nil {
				return nil, fmt.Errorf("failed to create database: %v", err)
			}
			dbConn, err = db.Connect(cfg)
			if err != nil {
				return nil, fmt.Errorf("failed to connect to database after creation: %v", err)
			}
		} else {
			return nil, fmt.Errorf("failed to connect to database: %v", err)
		}
	}

	// Run migrations
	err = dbConn.AutoMigrate(&model.Song{})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to run migrations")
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}
	logrus.Info("Migration completed: Table 'songs' has been created or already exists.")

	// Initialize the repository and service
	songRepository := repository.NewSongRepository(dbConn)
	songService := service.NewSongService(songRepository)

	// Set up the router
	r := gin.Default()
	router.SetupRoutes(r, songService) // Pass the service to the router
	return r, nil
}

func createDatabase(cfg *config.Config) error {
	// DSN for creating the database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	}
	defer db.Close()

	// Create the database
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", cfg.DBName))
	if err != nil {
		return fmt.Errorf("failed to create database %s: %v", cfg.DBName, err)
	}

	return nil
}

func TestIntegration_AddSong(t *testing.T) {
	r, err := setupIntegrationTest()
	assert.Nil(t, err, "Setting up integration test should not return an error")

	// Prepare request body
	reqBody := []byte(`{"group_name":"Muse","song_name":"Supermassive Black Hole"}`)
	req, _ := http.NewRequest("POST", "/api/v1/songs", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, http.StatusCreated, w.Code, "HTTP status should be 201")
	assert.Contains(t, w.Body.String(), "Supermassive Black Hole", "Response should contain the song name")
}
