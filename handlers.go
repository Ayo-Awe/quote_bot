package main

import (
	"fmt"

	"gopkg.in/telebot.v3"
)

func (a *Application) StartCommand(ctx telebot.Context) error {
	startMsg := fmt.Sprintf("Hello @%s 👋🏽\nWelcome to the quote bot!!!", ctx.Chat().Username)
	return ctx.Send(startMsg)
}

func (a *Application) QuoteCommand(ctx telebot.Context) error {
	quote, err := a.QuoteProvider.GetQuote("")
	if err != nil {
		return ctx.Send("An error occured while trying to load the quote ...")
	}

	quoteMsg := fmt.Sprintf("%s\n\n - %s", quote.Content, quote.Author)
	return ctx.Send(quoteMsg)
}
