package service

import (
	"context"

	"go.template/internal/models"
)

type HelloService interface {
	SayHello(ctx context.Context, req models.HelloRequest) (models.HelloResponse, error)
}
