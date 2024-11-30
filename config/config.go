package config

import (
	"path/filepath"
	"song-library/pkg/logger"

	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	APIBaseURL string
	ServerPort string
}

func LoadConfig(file string) (*Config, error) {
	dir, err := os.Getwd()
	if err != nil {
		logger.Error("Failed to get current directory", logrus.Fields{"error": err.Error()})
		return nil, err
	}
	envPath := filepath.Join(dir, file)

	if err := godotenv.Load(envPath); err != nil {
		logger.Error("Failed to load .env file", logrus.Fields{"path": envPath, "error": err.Error()})
		return nil, err
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		APIBaseURL: os.Getenv("API_BASE_URL"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}, nil
}
