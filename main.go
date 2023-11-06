package main

import (
	"link-bot/clients/telegram"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	tgClient = telegram.New(mustHost(), mustToken())

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	godotenv.Load()

	token := os.Getenv("BOT_TOKEN")

	if token == "" {
		log.Fatal("BOT_TOKEN is not set")
	}

	return token
}

func mustHost() string {
	godotenv.Load()

	host := os.Getenv("BOT_HOST")

	if host == "" {
		log.Fatal("BOT_HOST is not set")
	}

	return host
}
