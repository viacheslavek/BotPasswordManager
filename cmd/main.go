package main

import (
	"context"
	"fmt"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/clients/telegram"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/events/tgEvents"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/storage/postgresql"
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

	db, err := postgresql.New()
	if err != nil {
		log.Fatalf("can't up db %e", err)
	}

	err = db.Init(context.Background())
	if err != nil {
		log.Fatalf("can't init db %e", err)
	}

	fmt.Println(db)

	processor := tgEvents.NewProcessor(tgClient)

	log.Println("tgProcessor init")

	fmt.Println(processor)

	// fetcher = fetcher.New(tgClient)

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
