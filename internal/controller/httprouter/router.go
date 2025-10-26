package httprouter

import (
	"github.com/go-chi/chi/v5"
	"go.template/internal/controller/httprouter/middleware"
	"go.template/internal/handler"
	"go.template/internal/service/hello"
	"go.template/pkg/httpserver"
	zaplog "go.template/pkg/logger"
)

func NewRouter(server *httpserver.Server, logger *zaplog.Logger) {
	server.Mux.Use(middleware.LoggingMiddleware(logger))
	server.Mux.Use(middleware.RecoveryMiddleware(logger))
	helloSvc := hello.NewSvc(logger)
	helloHandler := handler.NewHelloHandler(helloSvc)
	server.Mux.Route("/v1", func(r chi.Router) {
		r.Get("/hello", helloHandler.SayHelloHandler()) // add handler here
	})
}
