package database

import (
	"database/sql"
	"fmt"
)

func InitDB() (*sql.DB, error) {
	connStr := "postgres://ppio:dummy@127.0.0.1/ppio?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db")
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("connection to DB is dead")
	}

	return db, nil
}
