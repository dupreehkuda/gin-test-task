package handlers

import (
	"github.com/rs/zerolog"

	i "github.com/dupreehkuda/gin-test-task/internal/interfaces"
)

// handlers provides access to service
type handlers struct {
	processor i.Processors
	logger    *zerolog.Logger
}

// New creates new instance of handlers
func New(processor i.Processors, logger *zerolog.Logger) *handlers {
	return &handlers{
		processor: processor,
		logger:    logger,
	}
}
