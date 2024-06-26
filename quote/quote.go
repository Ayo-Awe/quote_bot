package quote

import (
	"errors"
)

var (
	ErrQuoteNotFound       = errors.New("no matching quote found")
	ErrQuoteFetchFailed    = errors.New("couldn't fetch quote")
	ErrCategoryFetchFailed = errors.New("couldn't fetch categories")
	ErrQuoteSearchFailed   = errors.New("couldn't search for quote")
)

type QuoteParams struct {
	Limit    uint   `json:"limit"`
	Category string `json:"category"`
}

type QuoteProvider interface {
	GetQuote(category string) (*Quote, error)
	GetQuotes(params QuoteParams) ([]Quote, error)
	GetCategories() ([]Category, error)
	Search(query string) (*Quote, error)
}

type Quote struct {
	ID         string   `json:"id"`
	Content    string   `json:"content"`
	Author     string   `json:"author"`
	Categories []string `json:"categories"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
