package tgEvents

import (
	"log"
	"strings"
)

const ()

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command: '%s' from '%s'", text, username)

	// /set site login password  // по-хорошему тоже надо удалить, но не обязательно
	// /get site // после отправки нужно удалить сообщение
	// /del site //упс, надо переделать ручку удаления
	// /help
	// /start

	return nil
}
