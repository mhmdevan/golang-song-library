package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"song-library/internal/model"
	"song-library/internal/repository"
	"song-library/internal/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestHandler creates a test SongService and Gin engine
func setupTestHandler() (*service.SongService, *gin.Engine) {
	// Initialize in-memory database
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.Song{})

	// Create repository and service
	repo := repository.NewSongRepository(db)
	songService := service.NewSongService(repo)

	// Initialize Gin engine
	r := gin.Default()
	return songService, r
}

func TestAddSongHandler(t *testing.T) {
	// Setup test environment
	songService, r := setupTestHandler()

	// Register the handler
	r.POST("/songs", AddSong(songService))

	// Create a test request
	reqBody := []byte(`{"group_name":"Muse","song_name":"Supermassive Black Hole"}`)
	req, _ := http.NewRequest("POST", "/songs", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code, "HTTP status should be 201")
	assert.Contains(t, w.Body.String(), "Supermassive Black Hole", "Response should contain the song name")
}
