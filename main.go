package main

import (
	"encoding/json"
	"net/http"
	"quoteservice/quotes"

	"log"
)

func main() {
	qs, err := quotes.NewService()
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":8080", createHttpHandler(qs)))
}

func createHttpHandler(qs quotes.QuoteService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content/type", "application/json")

		q, err := json.Marshal(qs.Random())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(q))
	}
}
