package handler

import (
	"net/http"
	"song-library/internal/model"
	"song-library/internal/service"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

// SuccessResponse represents a standard success response
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// GetSongs retrieves all songs
// @Summary Retrieve all songs
// @Description Get a list of all songs
// @Tags songs
// @Produce json
// @Success 200 {object} SuccessResponse
// @Failure 500 {object} ErrorResponse
// @Router /songs [get]
func GetSongs(songService *service.SongService) gin.HandlerFunc {
	return func(c *gin.Context) {
		songs, err := songService.GetSongs()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error:   "Failed to retrieve songs",
				Details: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, SuccessResponse{Data: songs})
	}
}

// GetSongByID retrieves a song by its ID
// @Summary Retrieve a song by ID
// @Description Get a song by its unique ID
// @Tags songs
// @Produce json
// @Param id path string true "Song ID"
// @Success 200 {object} SuccessResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /songs/{id} [get]
func GetSongByID(songService *service.SongService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		song, err := songService.GetSongByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error:   "Failed to retrieve song",
				Details: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, SuccessResponse{Data: song})
	}
}

// AddSong adds a new song
// @Summary Add a new song
// @Description Create a new song with group name and song name
// @Tags songs
// @Accept json
// @Produce json
// @Param song body model.Song true "Song data"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /songs [post]
func AddSong(songService *service.SongService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var song model.Song
		if err := c.ShouldBindJSON(&song); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "Invalid request payload",
				Details: err.Error(),
			})
			return
		}

		if err := songService.AddSong(&song); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error:   "Failed to add song",
				Details: err.Error(),
			})
			return
		}
		c.JSON(http.StatusCreated, SuccessResponse{Data: song})
	}
}

// UpdateSong updates an existing song
// @Summary Update an existing song
// @Description Update a song's details by its ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Param song body model.Song true "Updated song data"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /songs/{id} [put]
func UpdateSong(songService *service.SongService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var song model.Song
		if err := c.ShouldBindJSON(&song); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Error:   "Invalid request payload",
				Details: err.Error(),
			})
			return
		}

		if err := songService.UpdateSong(id, &song); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error:   "Failed to update song",
				Details: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, SuccessResponse{Data: song})
	}
}

// DeleteSong deletes a song by its ID
// @Summary Delete a song
// @Description Remove a song from the database by its ID
// @Tags songs
// @Produce json
// @Param id path string true "Song ID"
// @Success 200 {object} SuccessResponse
// @Failure 500 {object} ErrorResponse
// @Router /songs/{id} [delete]
func DeleteSong(songService *service.SongService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := songService.DeleteSong(id); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Error:   "Failed to delete song",
				Details: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, SuccessResponse{Data: gin.H{"message": "Song deleted successfully"}})
	}
}
