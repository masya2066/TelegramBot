package main

import (
	tgClient "TelegramBot/clients/telegram"
	event_consumer "TelegramBot/consumer/event-consumer"
	telegram "TelegramBot/events/telegram"
	"TelegramBot/storage/files"
	"flag"
	"log"
)

// import "fmt"

const (
	host        = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventProcessor := telegram.New(
		tgClient.New(host, mustToken()),
		files.New(storagePath))

	log.Printf("Service is started!")

	if err := event_consumer.New(eventProcessor, eventProcessor, batchSize).Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("bot-token", "", "Usage telegram bot token")

	flag.Parse()

	if *token == "" {
		log.Fatal("Fatal Error")
	}

	return *token
}
