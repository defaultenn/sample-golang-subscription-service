package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

const (
	DefaultReadTimeout     = 5 * time.Second
	DefaultWriteTimeout    = 5 * time.Second
	DefaultAddr            = ":80"
	DefaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	Server          *http.Server
	NotifyChan      chan error
	ShutdownTimeout time.Duration
	logger          *zerolog.Logger
}

func New(ctx context.Context, handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  DefaultReadTimeout,
		WriteTimeout: DefaultWriteTimeout,
		Addr:         DefaultAddr,
	}

	l := zerolog.Ctx(ctx)

	s := &Server{
		Server:          httpServer,
		NotifyChan:      make(chan error, 1),
		ShutdownTimeout: DefaultShutdownTimeout,
		logger:          l,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.NotifyChan <- s.Server.ListenAndServe()
		close(s.NotifyChan)
	}()
	s.logger.Info().Msg(fmt.Sprintf("Start listening on %s", s.Server.Addr))
}

func (s *Server) Notify() <-chan error {
	return s.NotifyChan
}

func (s *Server) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, s.ShutdownTimeout)
	defer cancel()

	return s.Server.Shutdown(ctx)
}
