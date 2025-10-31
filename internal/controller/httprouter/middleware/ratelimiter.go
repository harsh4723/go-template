package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/httprate"
	zaplog "go.template/pkg/logger"
)

func RateLimitMiddlware(logger *zaplog.Logger, RequestLimMin int, opts ...httprate.Option) func(http.Handler) http.Handler {
	if len(opts) == 0 {
		return httprate.LimitByIP(RequestLimMin, time.Minute)
	}
	return httprate.Limit(RequestLimMin, time.Minute, opts...)
}
