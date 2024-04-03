package quote

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type quotableQuote struct {
	ID         string   `json:"_id"`
	Content    string   `json:"content"`
	Author     string   `json:"author"`
	AuthorSlug string   `json:"authorSlug"`
	Length     int      `json:"length"`
	Tags       []string `json:"tags"`
}

func (q *quotableQuote) ToQuote() Quote {
	return Quote{
		ID:         q.ID,
		Content:    q.Content,
		Author:     q.Author,
		Categories: q.Tags,
	}
}

type quotableProvider struct {
	BaseURL string
	client  *http.Client
}

func NewQuotableProvider() QuoteProvider {
	return &quotableProvider{
		client:  &http.Client{},
		BaseURL: "https://api.quotable.io",
	}
}

func (q *quotableProvider) GetQuote(category string) (*Quote, error) {
	quotes, err := q.GetQuotes(QuoteParams{Limit: 1, Category: category})

	if err != nil {
		return nil, err
	}

	if len(quotes) < 1 {
		return nil, ErrNoQuote
	}

	return &quotes[0], nil
}

func (q *quotableProvider) GetQuotes(params QuoteParams) ([]Quote, error) {

	url := fmt.Sprintf("%s/quotes/random?tags=%s&limit=%d", q.BaseURL, params.Category, params.Limit)

	res, err := q.client.Get(url)
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

	var quotables []quotableQuote
	if err = json.Unmarshal(resBody, &quotables); err != nil {
		return nil, err
	}

	var quotes []Quote
	for _, quotable := range quotables {
		quotes = append(quotes, quotable.ToQuote())
	}

	return quotes, nil
}

func (q *quotableProvider) GetCategories() ([]Category, error) {
	url := fmt.Sprintf("%s/tags", q.BaseURL)

	res, err := q.client.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, ErrCategoryFetchFailed
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var categories []Category
	err = json.Unmarshal(resBody, &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
