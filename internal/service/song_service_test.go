package service

import (
	"song-library/internal/model"
	"song-library/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestRepository creates an in-memory database for testing
func setupTestRepository() repository.SongRepository {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&model.Song{})
	return repository.NewSongRepository(db)
}

func TestSongService_AddSong(t *testing.T) {
	repo := setupTestRepository()
	songService := NewSongService(repo)

	song := &model.Song{
		GroupName: "Muse",
		SongName:  "Supermassive Black Hole",
	}

	err := songService.AddSong(song)
	assert.Nil(t, err, "Adding song should not return an error")

	songs, _ := songService.GetSongs()
	assert.Len(t, songs, 1, "There should be one song in the database")
	assert.Equal(t, "Supermassive Black Hole", songs[0].SongName, "Song name should match")
}

func TestSongService_GetSongs(t *testing.T) {
	repo := setupTestRepository()
	songService := NewSongService(repo)

	// Insert test data
	repo.AddSong(&model.Song{GroupName: "Muse", SongName: "Supermassive Black Hole"})

	songs, err := songService.GetSongs()
	assert.Nil(t, err, "Fetching songs should not return an error")
	assert.Len(t, songs, 1, "There should be one song in the database")
	assert.Equal(t, "Supermassive Black Hole", songs[0].SongName, "Song name should match")
}
