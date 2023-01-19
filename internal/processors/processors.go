package processors

import "github.com/rs/zerolog"

// processor provides service's business-logic
type processor struct {
	logger *zerolog.Logger
}

// New creates new instance of processor
func New(logger *zerolog.Logger) *processor {
	return &processor{
		logger: logger,
	}
}
