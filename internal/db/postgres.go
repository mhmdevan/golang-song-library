package db

import (
	"fmt"
	"song-library/config"
	"song-library/internal/model"
	"song-library/pkg/logger"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect initializes the database connection and applies migrations
func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to the database", logrus.Fields{
			"dsn":   dsn,
			"error": err.Error(),
		})
		return nil, errors.Wrap(err, "unable to connect to database")
	}

	// Run migrations
	err = db.AutoMigrate(&model.Song{})
	if err != nil {
		logger.Error("Failed to run migrations", logrus.Fields{"error": err.Error()})
		return nil, errors.Wrap(err, "migrations failed")
	}

	logger.Info("Database connected and migrations applied successfully", logrus.Fields{
		"host": cfg.DBHost,
		"port": cfg.DBPort,
	})
	return db, nil
}
