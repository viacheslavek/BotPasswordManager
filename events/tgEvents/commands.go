package tgEvents

import (
	"context"
	"fmt"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/storage"
	"log"
	"strings"
)

const (
	GetCmd   = "/get"
	AddCmd   = "/add"
	DelCmd   = "/del"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command: '%s' from '%s'", text, username)

	// /set site login password  // по-хорошему тоже надо удалить, но не обязательно
	// /get site // после отправки нужно удалить сообщение
	// /del site //упс, надо переделать ручку удаления

	switch text {
	case GetCmd:
		return p.addAccount(chatID, text, username)
	case AddCmd:
		return p.getAccount(chatID, text, username)
	case DelCmd:
		return p.delAccount(chatID, text, username)
	case HelpCmd:
		return p.help(chatID)
	case StartCmd:
		return p.start(chatID, username)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}

func (p *Processor) addAccount(chatID int, text string, username string) error {
	parseText := strings.Split(text, " ")
	// можно добавит каждому аргументу проверку на валидность,
	// а паролю проверку на устойчивость, но это на будущее
	if len(parseText) < 3 {
		err := p.help(chatID)
		if err != nil {
			return fmt.Errorf("can't sent help cmd %w\n", err)
		}
		return fmt.Errorf("not enouth arguments in add account\n")
	}

	acc := &storage.Account{
		Username: username,
		Site:     parseText[0],
		Login:    parseText[1],
		Password: parseText[2],
	}

	// можно добавить проверку на дубликат логина для сайта
	// дубликат пароля в общем
	// дубликат и логина, и пароля
	// и вообще надо все хэшировать
	// и по-хорошему солевой оператор добавить,
	// но, опять же, все потом

	err := p.storage.SaveAccount(context.Background(), acc)

	if err != nil {
		return fmt.Errorf("can't save account in db %w", err)
	}

	err = p.tg.SendMessage(chatID, msgAdd)
	if err != nil {
		return fmt.Errorf("can't send msg to user %w", err)
	}

	// по-хорошему добавить функцию удаления сообщения пользователя с паролем

	return nil
}

func (p *Processor) getAccount(chatID int, text string, username string) error {
	parseText := strings.Split(text, " ")
	if len(parseText) < 1 {
		err := p.help(chatID)
		if err != nil {
			return fmt.Errorf("can't sent help cmd %w\n", err)
		}
		return fmt.Errorf("not enouth arguments in get account\n")
	}

	accounts, err := p.storage.GetAccount(context.Background(), username, parseText[0])
	if err != nil {
		return fmt.Errorf("can't get account in db %w", err)
	}

	err = p.tg.SendMessage(chatID, msgGet)
	if err != nil {
		return fmt.Errorf("can't send msg to user %w", err)
	}

	for i, acc := range accounts {
		err = p.tg.SendMessage(chatID, fmt.Sprintf("%d:\nsite: %s\nlogin: %s\npassword: %s",
			i, acc.Site, acc.Login, acc.Password))
		if err != nil {
			return fmt.Errorf("can't send account msg to user %w\n", err)
		}
	}

	err = p.deleteMsgWithPassword(context.Background())
	if err != nil {
		return fmt.Errorf("can't delete msg to user with password %w\n", err)
	}

	return nil
}

func (p *Processor) delAccount(chatID int, text string, username string) error {
	parseText := strings.Split(text, " ")
	if len(parseText) < 1 {
		err := p.help(chatID)
		if err != nil {
			return fmt.Errorf("can't sent help cmd %w\n", err)
		}
		return fmt.Errorf("not enouth arguments in get account\n")
	}

	// у меня не так работает удаление в бд:
	// надо по фиксить, а пока так

	acc := &storage.Account{
		Username: username,
		Site:     parseText[0],
	}

	// было бы здорово забирать количество affected rows из бд и писать их здесь

	err := p.storage.DeleteAccount(context.Background(), acc)
	if err != nil {
		return fmt.Errorf("can't delete account in bd %w\n", err)
	}

	err = p.tg.SendMessage(chatID, msgDel)
	if err != nil {
		return fmt.Errorf("can't send msg to user %w", err)
	}

	return nil
}

func (p *Processor) help(chatID int) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) start(chatID int, username string) error {
	return p.tg.SendMessage(chatID, fmt.Sprintf("%s, %s", username, msgStart))
}

func (p *Processor) deleteMsgWithPassword(ctx context.Context) error {
	fmt.Println(ctx)
	log.Println("HEHHEHE\nя не удалю.")
	return nil
}
