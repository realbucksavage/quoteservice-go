package main

import (
	"net/http"
	"time"
)

type Middlware func(http.Handler) http.HandlerFunc

func loggingMiddlware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Infof("Request completed in %d ms", time.Since(start).Microseconds())
	}
}

func combineHandlers(root http.HandlerFunc, mwf ...Middlware) (handler http.HandlerFunc) {
	handler = root
	for _, m := range mwf {
		handler = m(handler)
	}
	return
}
