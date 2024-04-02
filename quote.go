package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	ErrNoQuote          = errors.New("no matching quote found")
	ErrQuoteFetchFailed = errors.New("couldn't fetch quote")
)

type Quote struct {
	ID         string   `json:"_id"`
	Content    string   `json:"content"`
	Author     string   `json:"author"`
	AuthorSlug string   `json:"authorSlug"`
	Length     int      `json:"length"`
	Tags       []string `json:"tags"`
}

type QuotableClient struct {
	URL
}

func (q *QuotableClient) GetQuote() (*Quote, error) {
	url := "https://api.quotable.io/quotes/random?tags=famous-quotes&limit=1"
	client := http.Client{}

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrQuoteFetchFailed
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var quotes []Quote
	if err = json.Unmarshal(resBody, &quotes); err != nil {
		return nil, err
	}

	if len(quotes) < 1 {
		return nil, ErrNoQuote
	}

	return &quotes[0], nil
}

func (q *QuotableClient) GetQuotes() ([]Quote, error) {

}
