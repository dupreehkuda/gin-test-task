package middleware

import "github.com/rs/zerolog"

// middleware provides services middleware
type middleware struct {
	logger *zerolog.Logger
}

// New creates new instance of middleware
func New(logger *zerolog.Logger) *middleware {
	return &middleware{
		logger: logger,
	}
}
