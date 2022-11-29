package main

import (
	"flag"
	"log"
	"secret-santa-bot/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	tgClient := telegram.New(mustToken(), tgBotHost)

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}

// mustToken позволяет получить токен из флага
func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token to access telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
