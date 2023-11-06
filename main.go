package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	t := mustToken()

	// tgClient = telegram.New(token)

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
