package tgEvents

import "github.com/VyacheslavIsWorkingNow/BotPasswordManager/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func NewProcessor(c telegram.Client) Processor {
	return Processor{}
}
