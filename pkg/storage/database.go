package storage

import (
	"database/sql"
	"log"
	"net/url"

	"github.com/pkg/errors"
)

func SetupDB(name string) (*sql.DB, func(), error) {
	q := make(url.Values)
	q.Set("sslmode", "disable")
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword("ppio", "dummy"),
		Host:     "127.0.0.1",
		Path:     name,
		RawQuery: q.Encode(),
	}

	db, err := sql.Open("postgres", u.String())
	if err != nil {
		return nil, nil, errors.Wrap(err, "open")
	}

	if err = db.Ping(); err != nil {
		return nil, nil, errors.Wrap(err, "ping")
	}

	tidy := func() {
		log.Println("closing db")
		db.Close()
	}

	return db, tidy, nil
}
