package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	t := mustToken()

	fmt.Println(t)

	log.Println("app starting")

	// tgClient = telegram.New(t)

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)

}

func mustToken() string {
	token := flag.String(
		"tg_token",
		"",
		"token for access telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("space token")
	}

	log.Println("get possible token")

	return *token
}
