package quotes

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
)

type QuoteService interface {
	Random() Quote
}

func (qs quoteSource) Random() Quote {
	idx := rand.Intn(len(qs.quotes))
	return qs.quotes[idx]
}

func NewService() (QuoteService, error) {
	b, err := ioutil.ReadFile("data/quotes.json")
	if err != nil {
		return nil, err
	}

	var q []Quote
	if err := json.Unmarshal(b, &q); err != nil {
		return nil, err
	}

	return quoteSource{q}, nil
}
