package main

import (
	"fmt"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/clients/telegram"
	"log"
	"os"
)

func main() {

	t := mustToken()

	fmt.Println(t)

	log.Println("app starting")

	tgClient := telegram.New("host", t)

	fmt.Println(tgClient)

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

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
