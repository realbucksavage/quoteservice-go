package main

import (
	"encoding/json"
	"net/http"
	"quoteservice/quotes"

	log "github.com/op/go-logging"
)

var (
	logger = log.MustGetLogger("quoteservice")
)

func main() {
	qs, err := quotes.NewService()
	if err != nil {
		logger.Fatal(err)
	}

	root := createHttpHandler(qs)
	handler := combineHandlers(root, rateLimiter, loggingMiddlware)

	logger.Info("Starting server on port 8080...")
	logger.Fatal(http.ListenAndServe(":8080", handler))
}

func createHttpHandler(qs quotes.QuoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		q, err := json.Marshal(qs.Random())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Errorf("marshal error : %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(q))
	}
}
