package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
)

// Config provides service address and paths to database
type Config struct {
	Address string `env:"SERVICE_ADDRESS" envDefault:":8080"`
}

// New creates new Config
func New(logger *zerolog.Logger) *Config {
	var config = Config{}
	var err = env.Parse(&config)
	if err != nil {
		logger.Error().Msg("Error occurred when parsing env")
	}

	flag.StringVar(&config.Address, "a", config.Address, "Launch address")
	flag.Parse()

	return &config
}
