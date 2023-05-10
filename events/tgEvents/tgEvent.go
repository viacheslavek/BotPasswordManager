package tgEvents

import (
	"errors"
	"fmt"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/clients/telegram"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/events"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/storage"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

type Meta struct {
	ChatID   int
	Username string
}

var (
	ErrUnknownEventType = errors.New("unknown type event")
	ErrUnknownMetaType  = errors.New("unknown meta type")
)

func NewProcessor(c *telegram.Client, s storage.Storage) *Processor {
	return &Processor{
		tg:      c,
		offset:  0,
		storage: s,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, fmt.Errorf("can't het Updates in Fetch %w\n", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		err := p.processMessage(event)
		if err != nil {
			return fmt.Errorf("can't process message %w\n", err)
		}
	default:
		return ErrUnknownEventType
	}

	return nil
}

func (p *Processor) processMessage(e events.Event) error {
	m, err := meta(e)
	if err != nil {
		return fmt.Errorf("can't process message %w\n", err)
	}

	if err = p.doCmd(e.Text, m.ChatID, m.Username); err != nil {
		return fmt.Errorf("can't do cmd in processMessage %w\n", err)
	}

	return nil
}

func meta(e events.Event) (Meta, error) {
	res, ok := e.Meta.(Meta)
	if !ok {
		return Meta{}, fmt.Errorf("not a meta %w\n", ErrUnknownMetaType)
	}
	return res, nil
}

func event(u telegram.Update) events.Event {
	res := events.Event{
		Type: fetchType(u),
		Text: fetchText(u),
	}

	if res.Type == events.Message {
		res.Meta = Meta{
			ChatID:   u.Message.Chat.ID,
			Username: u.Message.From.Username,
		}
	}
	return res
}

func fetchText(u telegram.Update) string {
	if u.Message == nil {
		return ""
	}
	return u.Message.Text
}

func fetchType(u telegram.Update) events.Type {
	if u.Message == nil {
		return events.Unknown
	} else {
		return events.Message
	}
}
