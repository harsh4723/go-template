package httprouter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
	"go.template/config"
	"go.template/internal/controller/httprouter/middleware"
	"go.template/internal/handler"
	"go.template/internal/service/hello"
	"go.template/pkg/httpserver"
	zaplog "go.template/pkg/logger"
)

func NewRouter(server *httpserver.Server, logger *zaplog.Logger, conf *config.Config) {
	server.Mux.Use(middleware.LoggingMiddleware(logger))
	server.Mux.Use(middleware.RecoveryMiddleware(logger))

	//go-chi ratelimiter (sliding window)
	//Rate-limit all routes at 100 req/min by IP address.
	server.Mux.Use(middleware.RateLimitMiddlware(logger, conf.RequestLimMin))

	helloSvc := hello.NewSvc(logger)

	helloHandler := handler.NewHelloHandler(helloSvc)
	server.Mux.Route("/v1", func(r chi.Router) {

		// Rate-limit v1 routes at 100 req/min by userID.
		r.Use(middleware.RateLimitMiddlware(logger, conf.RequestLimMin, httprate.WithKeyFuncs(func(r *http.Request) (string, error) {
			token, _ := r.Context().Value("userID").(string)
			return token, nil
		})))

		r.Get("/hello", helloHandler.SayHelloHandler()) // add handler here
	})
}
