package httpserver

import (
	"net"
	"strings"
	"time"
)

type Option func(*Server)

func ShortDuration(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}

func Port(port string) Option {
	return func(s *Server) {
		s.Server.Addr = net.JoinHostPort("", port)
		s.logger.Info().Msgf("HTTP Server Port: set %s", port)
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Server.ReadTimeout = timeout
		s.logger.Info().Msgf("HTTP Server Read Timeout: set %s", ShortDuration(timeout))
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Server.WriteTimeout = timeout
		s.logger.Info().Msgf("HTTP Server Write Timeout: set %s", ShortDuration(timeout))
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.ShutdownTimeout = timeout
		s.logger.Info().Msgf("HTTP Server Shutdown Timeout: set %s", ShortDuration(timeout))
	}
}
