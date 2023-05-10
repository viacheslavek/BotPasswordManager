package eventConsumer

import (
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/events"
	"log"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func NewConsumer(fetcher events.Fetcher, processor events.Processor, batchSize int) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c *Consumer) Start() error {
	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer %s", err.Error())
			continue

			// добавить правильную обработку ошибки
			// например, retried
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)

			continue
		}

		if err = c.handleEvents(gotEvents); err != nil {
			log.Println("err in handle event", err.Error())
		}
	}
}

func (c *Consumer) handleEvents(events []events.Event) error {
	for _, e := range events {
		log.Println("new event, do it")

		if err := c.processor.Process(e); err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
			// так же нужно обработать, но потом
			// так же: добавить горутины
			// так же: научиться не терять данные в случае ошибок (фоллбек)
			// так же: счетчик ошибок
		}
	}
	return nil
}
