package logger

import (
	"strings"

	"github.com/pkg/errors"
	log "go.uber.org/zap"
)

func level(level string) log.AtomicLevel {
	switch level {
	case "info":
		return log.NewAtomicLevelAt(log.InfoLevel)
	case "error":
		return log.NewAtomicLevelAt(log.ErrorLevel)
	case "debug":
		return log.NewAtomicLevelAt(log.DebugLevel)
	case "warn":
		return log.NewAtomicLevelAt(log.WarnLevel)
	default:
		return log.NewAtomicLevelAt(log.ErrorLevel)
	}
}

func NewLogger(logLevel string) (*log.Logger, error) {
	var (
		logger *log.Logger
		err    error
	)

	config := log.NewProductionConfig()
	config.Level = level(strings.ToLower(logLevel))

	logger, err = config.Build()
	if err != nil {
		return nil, errors.Wrap(err, "create logger from config failed")
	}

	return logger, err
}
