package db

import "database/sql"

func NewStorage(addr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", addr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
