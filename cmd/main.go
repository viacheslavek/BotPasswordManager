package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/clients/telegram"
	tgEvents "github.com/VyacheslavIsWorkingNow/BotPasswordManager/events/telegram"
	"log"
	"os"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {

	t := mustToken()

	log.Println("app starting")

	tgClient := telegram.NewClient(tgBotHost, t)

	log.Println("tgClient init")

	fmt.Println("ok all")

	// fetcher = fetcher.New(tgClient)

	processor := tgEvents.NewProcessor(tgClient)

	log.Println("tgProcessor init")

	fmt.Println(processor)

	// consumer.Start(fetcher, processor)

}

func mustToken() string {

	token := os.Getenv("TELEGRAM_TOKEN")

	if token == "" {
		log.Fatal("space token")
	}

	log.Println("get possible token")

	return token
}
