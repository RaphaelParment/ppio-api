package persistence

import (
	"github.com/jmoiron/sqlx"
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

func ConnectAndTidy(cfg *Config) (*sqlx.DB, func(l *log.Logger), error) {
	sslMode := "require"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host,
		Path:     cfg.Name,
		RawQuery: q.Encode(),
	}

	db, err := sqlx.Open("postgres", u.String())
	if err != nil {
		return nil, nil, errors.Wrap(err, "open")
	}

	if err = db.Ping(); err != nil {
		return nil, nil, errors.Wrap(err, "ping")
	}

	tidy := func(logger *log.Logger) {
		logger.Println("closing db")
		if err := db.Close(); err != nil {
			logger.Printf("failed to close db; %s", err)
		}
	}

	return db, tidy, nil
}
