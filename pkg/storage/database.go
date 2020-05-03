package storage

import (
	"database/sql"
	"log"
	"net/url"

	"github.com/pkg/errors"
)

type Config struct {
	User       string
	Password   string
	Host       string
	Name       string
	DisableTLS bool
}

func SetupDB(cfg *Config) (*sql.DB, func(), error) {
	sslMode := "require"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	// Query parameters.
	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")

	// Construct url.
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host,
		Path:     cfg.Name,
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
