package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"

	"github.com/dupreehkuda/gin-test-task/internal/config"
	"github.com/dupreehkuda/gin-test-task/internal/handlers"
	i "github.com/dupreehkuda/gin-test-task/internal/interfaces"
	"github.com/dupreehkuda/gin-test-task/internal/logger"
	"github.com/dupreehkuda/gin-test-task/internal/middleware"
	"github.com/dupreehkuda/gin-test-task/internal/processors"
)

type server struct {
	handlers i.Handlers
	mw       i.Middleware
	config   *config.Config
	logger   *zerolog.Logger
}

// NewByConfig returns server instance with default config
func NewByConfig() *server {
	log := logger.InitializeLogger()
	cfg := config.New(log)

	mv := middleware.New(log)

	logic := processors.New(log)

	handle := handlers.New(logic, log)

	return &server{
		handlers: handle,
		mw:       mv,
		config:   cfg,
		logger:   log,
	}
}

// Run runs the service
func (s server) Run() {
	srv := &http.Server{
		Addr:    s.config.Address,
		Handler: s.GetRouter(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal().Err(err).Msg("Error occurred starting server")
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	s.logger.Info().Msg("Server shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		s.logger.Fatal().Err(err).Msg("Shutdown error")
	}

	select {
	case <-ctx.Done():
		s.logger.Info().Msg("Timeout")
	}

	s.logger.Info().Msg("Successful shutdown")
}
