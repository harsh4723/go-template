package gotemp

import (
	"github.com/pkg/errors"
	"go.template/config"
	"go.template/internal/controller/httprouter"
	"go.template/pkg/httpserver"
	zaplog "go.template/pkg/logger"
)

type Option func(*gotemp) error

type gotemp struct {
	conf   *config.Config
	logger *zaplog.Logger
}

func New(opts ...Option) (*gotemp, error) {
	var (
		conf   *config.Config
		logger *zaplog.Logger
		err    error
	)

	conf = config.Load()
	logger, err = zaplog.NewLogger(conf.LogLevel)
	if err != nil {
		return nil, errors.Wrap(err, "create logger failed")
	}
	logger.Info("Intialized configs and logger")

	portOption := httpserver.Port(conf.Port)
	srv := httpserver.New(logger, portOption)
	httprouter.NewRouter(srv, logger)

	return &gotemp{conf: conf, logger: logger}, nil
}
