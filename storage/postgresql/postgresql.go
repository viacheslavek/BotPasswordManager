package postgresql

import (
	"context"
	"fmt"
	"github.com/VyacheslavIsWorkingNow/BotPasswordManager/storage"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

type Storage struct {
	db *pgx.Conn
}

func New() (*Storage, error) {
	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		return nil, fmt.Errorf("space db url\n")
	}

	db, err := pgx.Connect(context.Background(), databaseURL)

	// db, err := pgx.Connect(context.Background(), "postgres://slavabot:passwordbotpassword@localhost:5432/botpassword")

	if err != nil {
		return nil, fmt.Errorf("unable to connect database %w\n", err)
	}

	var greeting string
	err = db.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return nil, fmt.Errorf("QueryRow failed: %w\n", err)
	}

	log.Printf("connect db %s\n", greeting)

	return &Storage{db}, nil
}

func (s *Storage) Init(ctx context.Context) error {
	q := `
		CREATE TABLE IF NOT EXISTS client (
    		client_id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    		client_name VARCHAR(255)
		);

		CREATE TABLE IF NOT EXISTS site (
  		  site_id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  		  site_name VARCHAR(255)
		);

		CREATE TABLE IF NOT EXISTS account (
   			account_id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
   			client_id BIGINT,
  			site_id BIGINT,
    		password VARCHAR(255),
   			login VARCHAR(255),
    		CONSTRAINT "FK_client_id"
        		FOREIGN KEY (client_id) REFERENCES client (client_id),
    		CONSTRAINT "FK_site_id"
       			FOREIGN KEY (site_id) REFERENCES site (site_id)
		    );
	`

	_, err := s.db.Exec(
		ctx,
		q,
	)
	if err != nil {
		return fmt.Errorf("can't create tables %w", err)
	}

	return nil
}

func (s *Storage) SaveAccount(ctx context.Context, a *storage.Account) error {

	// по-хорошему надо сделать одну функцию извлечения сайта и клиента
	// и передавать название таблицы и атрибуты как параметры

	clientId, err := s.addClient(ctx, a.Username)
	if err != nil {
		return fmt.Errorf("can't find or add client %w", err)
	}

	siteId, err := s.addSite(ctx, a.Site)
	if err != nil {
		return fmt.Errorf("can't find or add site %w", err)
	}

	// Пока сделаю мвп: в случае совпадения логина и пароля я просто его перезатру:
	// В случае совпадения логина у меня будет два пароля на один логин
	// Если успею, исправлю это, но потом

	err = s.addAccount(ctx, a, clientId, siteId)

	if err != nil {
		return fmt.Errorf("can't add account %w", err)
	}

	return nil
}

func (s *Storage) GetAccount(ctx context.Context, username string, site string) ([]storage.Account, error) {

	q := `
		SELECT login, password
		FROM account
   			INNER JOIN client c on account.client_id = c.client_id AND 
   			                     c.client_name = $1
		    INNER JOIN site s on account.site_id = s.site_id AND
		                         s.site_name = $2;
	`

	rows, err := s.db.Query(ctx, q, username, site)
	defer rows.Close()

	if err != nil {
		return nil, fmt.Errorf("can't select account %w\n", err)
	}

	ans := make([]storage.Account, 0)

	for rows.Next() {
		var temp storage.Account
		err = rows.Scan(&temp.Login, &temp.Password)
		if err != nil {
			return nil, fmt.Errorf("can't scan account %w\n", err)
		}
		ans = append(ans, temp)
	}

	return ans, nil
}

func (s *Storage) DeleteAccount(ctx context.Context, a *storage.Account) error {

	siteId, err := s.addSite(ctx, a.Site)
	if err != nil {
		return fmt.Errorf("can't add site in delete %w\n", err)
	}
	clientId, err := s.addClient(ctx, a.Username)
	if err != nil {
		return fmt.Errorf("can't add client in delete %w\n", err)
	}

	q := `
		DELETE FROM account ()
		WHERE site_id = $1 AND client_id = $2 AND password = $3 AND login = $4
	`

	_, err = s.db.Query(ctx, q, siteId, clientId, a.Password, a.Login)

	if err != nil {
		return fmt.Errorf("can't delete account %w\n", err)
	}

	return nil
}

func (s *Storage) addClient(ctx context.Context, client string) (int, error) {
	q := `
		SELECT count(*), client_id
		FROM client
		WHERE client_name = $1;
	`

	rows, err := s.db.Query(ctx, q, client)
	defer rows.Close()

	if err != nil {
		return 0, fmt.Errorf("can't select client %w\n", err)
	}

	var count int
	var clientId int
	err = rows.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("can't scan client %w\n", err)
	}

	if count == 0 {
		// add user in table
		q = `
			INSERT INTO client (client_name)
			VALUES (
        		$1
			);
		`
		_, err = s.db.Exec(ctx, q, client)
		if err != nil {
			return 0, fmt.Errorf("can't insert client %w\n", err)
		}

		q = `
			SELECT client_id
			FROM client
			WHERE client_name = $1;
		`

		rows2, errIf := s.db.Query(ctx, q, client)
		defer rows2.Close()

		if errIf != nil {
			return -1, fmt.Errorf("can't select client where insert client %w\n", err)
		}

		err = rows.Scan(&clientId)
		if err != nil {
			return -1, fmt.Errorf("can't scan client where insert client %w\n", err)
		}
	} else {
		err = rows.Scan(&clientId)
		if err != nil {
			return -1, fmt.Errorf("can't scan site %w\n", err)
		}
	}

	return clientId, nil
}

func (s *Storage) addSite(ctx context.Context, site string) (int, error) {
	q := `
		SELECT count(*), site_id
		FROM site
		WHERE site_name = $1;
	`

	rows, err := s.db.Query(ctx, q, site)
	defer rows.Close()

	if err != nil {
		return -1, fmt.Errorf("can't select site %w\n", err)
	}

	var count int
	var siteId int
	err = rows.Scan(&count)
	if err != nil {
		return -1, fmt.Errorf("can't scan site %w\n", err)
	}

	if count == 0 {
		// add site in table
		q = `
			INSERT INTO site (site_name)
			VALUES (
        		'$1'
			);
		`
		_, err = s.db.Exec(ctx, q, site)
		if err != nil {
			return -1, fmt.Errorf("can't insert site %w\n", err)
		}

		q = `
			SELECT site_id
			FROM site
			WHERE site_name = $1;
		`

		rows2, errIf := s.db.Query(ctx, q, site)
		defer rows2.Close()

		if errIf != nil {
			return -1, fmt.Errorf("can't select site where insert site %w\n", err)
		}

		err = rows.Scan(&siteId)
		if err != nil {
			return -1, fmt.Errorf("can't scan site where insert site %w\n", err)
		}

	} else {
		err = rows.Scan(&siteId)
		if err != nil {
			return -1, fmt.Errorf("can't scan site %w\n", err)
		}
	}

	return siteId, nil
}

func (s *Storage) addAccount(ctx context.Context, a *storage.Account, clientId, siteId int) error {

	// кидаю запрос на поиск
	q := `
		SELECT count(*)
		FROM account
		WHERE client_id = $1 AND site_id = $2;
	`

	rows, err := s.db.Query(ctx, q, clientId, siteId)
	defer rows.Close()

	if err != nil {
		return fmt.Errorf("can't select account %w\n", err)
	}

	var count int
	err = rows.Scan(&count)
	if err != nil {
		return fmt.Errorf("can't scan account %w\n", err)
	}

	// если нет, добавляю
	// если есть, обновляю

	if count == 0 {
		q = `
			INSERT INTO account (client_id, site_id, password, login)
			VALUES (
			        $1, $2, $3, $4
			);
		`
		_, err = s.db.Exec(ctx, q, clientId, siteId, a.Password, a.Login)
		if err != nil {
			return fmt.Errorf("can't insert account %w\n", err)
		}

	} else {
		q = `
			UPDATE account
			SET login = $4
			WHERE client_id = $1 AND site_id = $2;

			UPDATE account
			SET password = $3
			WHERE client_id = $1 AND site_id = $2;
		`
		_, err = s.db.Exec(ctx, q, clientId, siteId, a.Password, a.Login)
		if err != nil {
			return fmt.Errorf("can't update account %w\n", err)
		}
	}

	return nil
}
