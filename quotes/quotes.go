package quotes

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
)

type Service interface {
	Random() Quote
}

func (qs quoteSource) Random() Quote {
	idx := rand.Intn(len(qs.quotes))
	return qs.quotes[idx]
}

func NewService() (Service, error) {
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
