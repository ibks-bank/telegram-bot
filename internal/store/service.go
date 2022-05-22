package store

import (
	"database/sql"
	"github.com/lib/pq"

	"github.com/ibks-bank/libs/cerr"
)

type store struct {
	db *sql.DB
}

func New(db *sql.DB) *store {
	return &store{db: db}
}

func (s *store) InsertUser(username string) error {
	_, err := s.db.Exec("insert into users(username) values ($1)", username)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return nil
		}
		return cerr.Wrap(err, "can't exec context")
	}

	return nil
}

//InsertToken(ctx context.Context, username, token string) error
//GetToken(ctx context.Context, username string) (string, error)

func (s *store) UpdateToken(username, token string) error {
	_, err := s.db.Exec("update users set token = $1 where username = $2", token, username)
	if err != nil {
		return cerr.Wrap(err, "can't exec context")
	}

	return nil
}

func (s *store) GetToken(username string) (string, error) {
	row := s.db.QueryRow("select token from users where username = $1", username)

	token := ""
	err := row.Scan(&token)
	if err != nil {
		return "", cerr.Wrap(err, "can't scan row")
	}

	return token, nil
}
