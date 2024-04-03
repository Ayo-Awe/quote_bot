package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/ayo-awe/quote_bot/quote"
	"gopkg.in/telebot.v3"
)

func (a *Application) StartCommand(ctx telebot.Context) error {
	startMsg := fmt.Sprintf("Hello @%s 👋🏽\nWelcome to the quote bot!!!", ctx.Chat().Username)
	return ctx.Send(startMsg)
}

func (a *Application) QuoteCommand(ctx telebot.Context) error {
	category := strings.TrimSpace(ctx.Message().Payload)

	q, err := a.QuoteProvider.GetQuote(category)
	if err != nil {
		errMsg := "An error occured while trying to load quote ..."

		if errors.Is(err, quote.ErrQuoteNotFound) {
			errMsg = "I couldn't find a quote that matches the specified category"
		}

		return ctx.Send(errMsg)
	}

	quoteMsg := fmt.Sprintf("%s\n\n - %s", q.Content, q.Author)
	return ctx.Send(quoteMsg)
}

func (a *Application) ListCategoriesCommand(ctx telebot.Context) error {
	categories, err := a.QuoteProvider.GetCategories()
	if err != nil {
		return ctx.Send("An error occured while trying to load categories...")
	}

	tmp, err := template.New("categories").Parse("Here's a list of all available categories\n\n{{range .}}- {{.Name}}\n{{end}}")
	if err != nil {
		return ctx.Send("An error occured while trying to load categories...")
	}

	buf := bytes.Buffer{}
	if err = tmp.Execute(&buf, categories); err != nil {
		return ctx.Send("An error occured while trying to load categories...")
	}

	return ctx.Send(buf.String())
}

func (a *Application) SearchCommand(ctx telebot.Context) error {
	q, err := a.QuoteProvider.Search(ctx.Message().Payload)

	if err != nil {
		errMsg := "An error occured while trying to search..."
		if errors.Is(err, quote.ErrQuoteNotFound) {
			errMsg = "No matching quote found"
		}
		return ctx.Send(errMsg)
	}

	if err = ctx.Send("Found a matching quote"); err != nil {
		return err
	}

	quoteMsg := fmt.Sprintf("%s\n\n - %s", q.Content, q.Author)
	return ctx.Send(quoteMsg)
}
