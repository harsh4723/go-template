package gotemp

import (
	"os"
	"os/signal"

	"github.com/pkg/errors"
	"go.template/config"
	"go.template/internal/controller/httprouter"
	"go.template/pkg/httpserver"
	zaplog "go.template/pkg/logger"
	"go.uber.org/zap"
)

type Option func(*gotemp) error

type gotemp struct {
	conf   *config.Config
	logger *zaplog.Logger
	server *httpserver.Server
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
	httprouter.NewRouter(srv, logger, conf)

	return &gotemp{conf: conf, logger: logger, server: srv}, nil
}

func (gt *gotemp) Open() error {

	go func(s *httpserver.Server) {
		gt.logger.Info(
			"Starting Server at Addr",
			zap.String("Name", gt.conf.AppName),
			zap.String("Port", gt.conf.Port),
		)

		err := gt.server.Open()
		if err != nil {
			gt.logger.Panic(err.Error())
		}
	}(gt.server)

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c
	gt.logger.Info("Recieved SIGINT, shutting down server")
	err := gt.server.Close()
	if err != nil {
		gt.logger.Panic("Error Closing Server: ", zap.Error(err))
	}
	os.Exit(0)
	return nil
}
