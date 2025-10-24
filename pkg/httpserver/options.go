package httpserver

import (
	"net"
	"time"
)

// Option -.
type Option func(*Server)

// Port -.
func Port(port string) Option {
	return func(s *Server) {
		s.Addr = net.JoinHostPort("", port)
	}
}

// ReadTimeout -.
func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.ReadTimeout = timeout
	}
}

// WriteTimeout -.
func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.WriteTimeout = timeout
	}
}

// IdleTimeout -.
func IdleTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.IdleTimeout = timeout
	}
}
