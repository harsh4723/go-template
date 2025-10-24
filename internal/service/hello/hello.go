package hello

import (
	"context"

	"go.template/internal/models"
	"go.template/internal/service"
	zaplog "go.template/pkg/logger"
)

type svc struct {
	logger zaplog.Logger
}

func (s *svc) SayHello(ctx context.Context, req models.HelloRequest) (models.HelloResponse, error) {
	return models.HelloResponse{}, nil
}

func New(logger zaplog.Logger) service.HelloService {
	return &svc{logger: logger}
}
