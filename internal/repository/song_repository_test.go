package repository

import (
	"song-library/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestRepository creates an in-memory database and returns a SongRepository
func setupTestRepository() SongRepository {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.Song{})
	return NewSongRepository(db)
}

func TestSongRepository_AddSong(t *testing.T) {
	repo := setupTestRepository()

	song := &model.Song{
		GroupName: "Muse",
		SongName:  "Supermassive Black Hole",
	}

	err := repo.AddSong(song)
	assert.Nil(t, err, "Adding song should not return an error")

	// Verify the song was added
	songs, _ := repo.GetSongs()
	assert.Len(t, songs, 1, "There should be one song in the database")
	assert.Equal(t, "Supermassive Black Hole", songs[0].SongName, "Song name should match")
}

func TestSongRepository_GetSongs(t *testing.T) {
	repo := setupTestRepository()

	// Add test data
	repo.AddSong(&model.Song{GroupName: "Muse", SongName: "Supermassive Black Hole"})

	// Fetch songs
	songs, err := repo.GetSongs()
	assert.Nil(t, err, "Fetching songs should not return an error")
	assert.Len(t, songs, 1, "There should be one song in the database")
	assert.Equal(t, "Supermassive Black Hole", songs[0].SongName, "Song name should match")
}

func TestSongRepository_GetSongByID(t *testing.T) {
	repo := setupTestRepository()

	// Add test data
	song := &model.Song{GroupName: "Muse", SongName: "Supermassive Black Hole"}
	repo.AddSong(song)

	// Fetch song by ID
	result, err := repo.GetSongByID("1")
	assert.Nil(t, err, "Fetching song by ID should not return an error")
	assert.Equal(t, "Supermassive Black Hole", result.SongName, "Song name should match")
}

func TestSongRepository_DeleteSong(t *testing.T) {
	repo := setupTestRepository()

	// Add test data
	song := &model.Song{GroupName: "Muse", SongName: "Supermassive Black Hole"}
	repo.AddSong(song)

	// Delete the song
	err := repo.DeleteSong("1")
	assert.Nil(t, err, "Deleting song should not return an error")

	// Verify the song was deleted
	songs, _ := repo.GetSongs()
	assert.Len(t, songs, 0, "The database should be empty after deletion")
}
