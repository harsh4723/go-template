package httpserver

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	zaplog "go.template/pkg/logger"
)

const (
	_defaultAddr         = ":8080"
	_defaultReadTimeout  = 5 * time.Second
	_defaultWriteTimeout = 5 * time.Second
	_defaultIdleTimeout  = 60 * time.Second
)

type Server struct {
	*http.Server
	Mux *chi.Mux
	log *zaplog.Logger
}

func New(logger *zaplog.Logger, opts ...Option) *Server {
	mux := chi.NewRouter()
	server := &Server{
		&http.Server{
			Addr:         _defaultAddr,
			IdleTimeout:  _defaultIdleTimeout,
			WriteTimeout: _defaultWriteTimeout,
			ReadTimeout:  _defaultReadTimeout,
		},
		mux,
		logger,
	}

	for _, opt := range opts {
		opt(server)
	}

	// bind mux to http server
	server.Handler = mux

	return server
}
