package tgEvents

import (
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/clients/telegram"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/storage"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

func NewProcessor(c *telegram.Client, s storage.Storage) *Processor {
	return &Processor{
		tg:      c,
		offset:  0,
		storage: s,
	}
}

func (p *Processor) Fetch() {

}
