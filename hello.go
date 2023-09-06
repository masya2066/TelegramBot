package main

import (
	"flag"
	"log"
)

// import "fmt"

const (
	host = "api.telegram.org"
)

func main() {
	tgCLient := telegram.New(host, mustToken())
}

func mustToken() string {
	token := flag.String("bot-token", "", "Usage telegram bot token")

	flag.Parse()

	if *token == "" {
		log.Fatal("Fatal Error")
	}

	return *token
}
