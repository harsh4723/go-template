package gotemp

import (
	"github.com/pkg/errors"
	"go.template/config"
	zaplog "go.template/pkg/logger"
)

func WithConfOption() Option {
	return func(g *gotemp) error {
		conf := config.Load()
		if conf == nil {
			return errors.New("configs not initialized properly")
		}
		g.conf = conf
		return nil
	}
}

func WithLoggerOption() Option {
	return func(g *gotemp) error {
		var (
			logger *zaplog.Logger
			err    error
		)
		logger, err = zaplog.NewLogger(g.conf.LogLevel)
		if err != nil {
			return errors.Wrap(err, "create logger failed")
		}
		logger.Info("Intialized logger")
		g.logger = logger
		return nil
	}
}
