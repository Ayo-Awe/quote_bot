package main

import (
	"flag"
	"fmt"
	"log"

	"gopkg.in/telebot.v3"
)

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

	b.Handle("/start", func(ctx telebot.Context) error {
		startMsg := fmt.Sprintf("Hello @%s 👋🏽\nWelcome to the quote bot!!!", ctx.Chat().Username)
		return ctx.Send(startMsg)
	})

	fmt.Print("Starting quote bot...")
	b.Start()
}
