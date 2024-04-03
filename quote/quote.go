package quote

import (
	"errors"
)

var (
	ErrNoQuote          = errors.New("no matching quote found")
	ErrQuoteFetchFailed = errors.New("couldn't fetch quote")
)

type QuoteParams struct {
	Limit    uint   `json:"limit"`
	Category string `json:"category"`
}

type QuoteProvider interface {
	GetQuote(category string) (*Quote, error)
	GetQuotes(params QuoteParams) ([]Quote, error)
}

type Quote struct {
	ID         string   `json:"id"`
	Content    string   `json:"content"`
	Author     string   `json:"author"`
	Categories []string `json:"categories"`
}
