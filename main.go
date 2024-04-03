package main

import (
	"flag"
	"fmt"
	"go-telebot/quote"
	"log"

	"gopkg.in/telebot.v3"
)

type Application struct {
	QuoteProvider quote.QuoteProvider
}

func main() {

	port := flag.Uint("port", 4000, "An available system port")
	webhookUrl := flag.String("webhook_url", "https://webhook.site", "A public url")
	token := flag.String("token", "", "A telegram bot token")

	flag.Parse()

	poller := &telebot.Webhook{
		Endpoint: &telebot.WebhookEndpoint{
			PublicURL: *webhookUrl,
		},
		Listen: fmt.Sprintf(":%d", *port),
	}

	teleConfig := telebot.Settings{
		Token:  *token,
		Poller: poller,
	}

	b, err := telebot.NewBot(teleConfig)
	if err != nil {
		log.Fatal(err)
	}

	app := &Application{QuoteProvider: quote.NewQuotableProvider()}

	b.Handle("/start", app.StartCommand)
	b.Handle("/quote", app.QuoteCommand)

	fmt.Print("Starting quote bot...")
	b.Start()
}
