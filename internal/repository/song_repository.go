package repository

import (
	"song-library/internal/model"

	"gorm.io/gorm"
)

// SongRepository defines methods for interacting with the songs database
type SongRepository interface {
	GetSongs() ([]model.Song, error)
	GetSongByID(id string) (*model.Song, error)
	AddSong(song *model.Song) error
	UpdateSong(id string, song *model.Song) error
	DeleteSong(id string) error
}

// songRepository implements SongRepository
type songRepository struct {
	db *gorm.DB
}

// NewSongRepository creates a new SongRepository
func NewSongRepository(db *gorm.DB) SongRepository {
	return &songRepository{db: db}
}

func (r *songRepository) GetSongs() ([]model.Song, error) {
	var songs []model.Song
	if err := r.db.Find(&songs).Error; err != nil {
		return nil, err
	}
	return songs, nil
}

func (r *songRepository) GetSongByID(id string) (*model.Song, error) {
	var song model.Song
	if err := r.db.First(&song, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &song, nil
}

func (r *songRepository) AddSong(song *model.Song) error {
	return r.db.Create(song).Error
}

func (r *songRepository) UpdateSong(id string, song *model.Song) error {
	return r.db.Model(&model.Song{}).Where("id = ?", id).Updates(song).Error
}

func (r *songRepository) DeleteSong(id string) error {
	return r.db.Delete(&model.Song{}, "id = ?", id).Error
}
