package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// InitializeLogger initializes new logger instance
func InitializeLogger() *zerolog.Logger {
	logger := zerolog.New(os.Stdout)
	return &logger
}
