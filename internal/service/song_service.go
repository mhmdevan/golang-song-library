package service

import (
	"song-library/internal/model"
	"song-library/internal/repository"
)

type SongService struct {
	repo repository.SongRepository
}

func NewSongService(repo repository.SongRepository) *SongService {
	return &SongService{repo: repo}
}

func (s *SongService) GetSongs() ([]model.Song, error) {
	return s.repo.GetSongs()
}

func (s *SongService) GetSongByID(id string) (*model.Song, error) {
	return s.repo.GetSongByID(id)
}

func (s *SongService) AddSong(song *model.Song) error {
	return s.repo.AddSong(song)
}

func (s *SongService) UpdateSong(id string, song *model.Song) error {
	return s.repo.UpdateSong(id, song)
}

func (s *SongService) DeleteSong(id string) error {
	return s.repo.DeleteSong(id)
}
