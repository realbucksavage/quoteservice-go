package main

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

type Middlware func(http.Handler) http.HandlerFunc

func loggingMiddlware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Infof("Request completed in %d ms", time.Since(start).Microseconds())
	}
}

func rateLimiter(next http.Handler) http.HandlerFunc {

	limiter := rate.NewLimiter(1, 1)

	return func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)

			e := http.StatusText(http.StatusTooManyRequests)
			w.Write([]byte(e))

			logger.Warningf("Request from %s dropped (rate-limited)", r.RemoteAddr)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func combineHandlers(root http.HandlerFunc, mwf ...Middlware) (handler http.HandlerFunc) {
	handler = root
	for _, m := range mwf {
		handler = m(handler)
	}
	return
}
