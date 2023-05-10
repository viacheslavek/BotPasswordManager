package main

import (
	"context"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/clients/telegram"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/comsumer/eventConsumer"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/events/tgEvents"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/storage/postgresql"
	"log"
	"os"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {

	t := mustToken()

	log.Println("app starting")

	tgClient := telegram.NewClient(tgBotHost, t)

	log.Println("tgClient init")

	// time.Sleep(1 * time.Second)

	db, err := postgresql.New()
	if err != nil {
		log.Fatalf("can't up db %e", err)
	}

	err = db.Init(context.Background())
	if err != nil {
		log.Fatalf("can't init db %e", err)
	}

	eventsProcessor := tgEvents.NewProcessor(&tgClient, db)

	log.Println("eventsProcessor init")

	consumer := eventConsumer.NewConsumer(eventsProcessor, eventsProcessor, batchSize)

	if err = consumer.Start(); err != nil {
		log.Fatalf("consumer dead :( %e", err)
	}

}

func mustToken() string {

	token := os.Getenv("TELEGRAM_TOKEN")

	if token == "" {
		log.Fatal("space token")
	}

	log.Println("get possible token")

	return token
}
