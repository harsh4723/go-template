package hello

import (
	"context"
	"fmt"

	"go.template/internal/models"
	"go.template/internal/service"
	zaplog "go.template/pkg/logger"
)

type svc struct {
	logger *zaplog.Logger
}

func (s *svc) SayHello(ctx context.Context, req models.HelloRequest) (models.HelloResponse, error) {
	msg := fmt.Sprintf("Hello %s", req.Name)

	return models.HelloResponse{Message: msg}, nil
}

func NewSvc(logger *zaplog.Logger) service.HelloService {
	return &svc{logger: logger}
}
