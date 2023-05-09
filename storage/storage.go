package storage

type Storage interface {
	SaveAccount(a *Account) error
	GetAccount(username string, site string) (Account, error)
	DeleteAccount(a *Account) error
}

type Account struct {
	Username string
	Site     string
	Login    string
	Password string
}
