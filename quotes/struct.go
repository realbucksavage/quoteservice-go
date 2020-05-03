package quotes

type quoteSource struct {
	quotes []Quote
}

type Quote struct {
	Content string   `json:"content"`
	Author  string   `json:"author"`
	Tags    []string `json:"tags"`
}
