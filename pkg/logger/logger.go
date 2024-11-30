package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Fields is an alias for logrus.Fields to simplify usage
type Fields = logrus.Fields

var log = logrus.New()

func Init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

// Info logs informational messages
func Info(message string, fields Fields) {
	log.WithFields(fields).Info(message)
}

// Error logs error messages
func Error(message string, fields Fields) {
	log.WithFields(fields).Error(message)
}

// Debug logs debug messages
func Debug(message string, fields Fields) {
	log.WithFields(fields).Debug(message)
}
