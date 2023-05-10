package storage

import "context"

type Storage interface {
	SaveAccount(ctx context.Context, a *Account) error
	GetAccount(ctx context.Context, username string, site string) ([]Account, error)
	DeleteAccount(ctx context.Context, a *Account) error
}

type Account struct {
	Username string
	Site     string
	Login    string
	Password string
}
