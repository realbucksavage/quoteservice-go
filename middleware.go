package main

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Middleware func(http.Handler) http.Handler

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Infof("Request completed in %d ms", time.Since(start).Microseconds())
	})
}

func rateLimiter(limit int) Middleware {

	return func(next http.Handler) http.Handler {
		limiter := rate.NewLimiter(1, 200)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				w.WriteHeader(http.StatusTooManyRequests)

				e := http.StatusText(http.StatusTooManyRequests)
				w.Write([]byte(e))

				logger.Warningf("Request from %s dropped (rate-limited)", r.RemoteAddr)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func combineHandlers(root http.Handler, mwf ...Middleware) (handler http.Handler) {
	handler = root
	for _, m := range mwf {
		handler = m(handler)
	}
	return
}
