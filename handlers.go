package main

import (
	"errors"
	"fmt"
	"strings"

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

		if errors.Is(err, quote.ErrNoQuote) {
			errMsg = "I couldn't find a quote that matches the specified category"
		}

		return ctx.Send(errMsg)
	}

	quoteMsg := fmt.Sprintf("%s\n\n - %s", q.Content, q.Author)
	return ctx.Send(quoteMsg)
}
