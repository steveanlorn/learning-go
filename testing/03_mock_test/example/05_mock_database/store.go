package store

import (
	"context"
	"database/sql"
)

const queryInsertUser = `
	INSERT INTO
		user
	(
		username,
		password
	) VALUES (
		@username,
		@password
	) RETURNING user_id
`

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	store := Store{
		db: db,
	}

	return &store
}

func (s *Store) CreateUser(ctx context.Context, username string, password string) (int, error) {
	var err error
	if err = s.db.PingContext(ctx); err != nil {
		return -1, err
	}

	tx, err := s.db.Begin()
	if err != nil {
		return -1, err
	}

	stmt, err := tx.PrepareContext(ctx, queryInsertUser)
	if err != nil {
		return -1, err
	}

	defer func() {
		stmtCloseErr := stmt.Close()
		if err == nil {
			err = stmtCloseErr
		}
	}()

	row := stmt.QueryRowContext(ctx,
		sql.Named("username", username),
		sql.Named("password", password),
	)

	var userID int
	if err = row.Scan(&userID); err != nil {
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		return -1, err
	}

	return userID, nil
}
