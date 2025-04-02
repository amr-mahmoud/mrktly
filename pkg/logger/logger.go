package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func NewLogger(env string) *Logger {
	// Configure console writer
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	// Create base logger
	log := zerolog.New(output).With().Timestamp().Logger()

	// Set global log level based on environment
	switch env {
	case "production":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "test":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	default: // development
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return &Logger{log}
}