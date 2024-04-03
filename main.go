package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/ayo-awe/quote_bot/quote"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

type Application struct {
	QuoteProvider quote.QuoteProvider
}

func main() {

	port := flag.Uint("port", 4000, "An available system port")
	webhookUrl := flag.String("webhook_url", "https://webhook.site", "A public url")
	token := flag.String("token", "", "A telegram bot token")

	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)
	publicUrl := fmt.Sprintf("%s/%s", strings.TrimSuffix(*webhookUrl, "/"), *token)

	poller := &telebot.Webhook{
		Endpoint: &telebot.WebhookEndpoint{
			PublicURL: publicUrl,
		},
		Listen: addr,
	}

	teleConfig := telebot.Settings{
		Token:  *token,
		Poller: poller,
	}

	bot, err := telebot.NewBot(teleConfig)
	if err != nil {
		log.Fatal(err)
	}

	bot.Use(middleware.Logger())

	app := &Application{QuoteProvider: quote.NewQuotableProvider()}

	bot.Handle("/start", app.StartCommand)
	bot.Handle("/quote", app.QuoteCommand)

	fmt.Print("Starting quote bot...")
	bot.Start()
}
