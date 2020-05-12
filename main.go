package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	ppioHTTP "github.com/RaphaelParment/ppio-api/pkg/http"
	"github.com/RaphaelParment/ppio-api/pkg/storage"
	"github.com/gorilla/handlers"
	"github.com/pkg/errors"

	"github.com/ardanlabs/conf"
	_ "github.com/lib/pq"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	var cfg struct {
		DB struct {
			User       string `conf:"default:ppio"`
			Password   string `conf:"default:dummy,noprint"`
			Host       string `conf:"default:0.0.0.0"`
			Name       string `conf:"default:ppio"`
			DisableTLS bool   `conf:"default:false"`
		}
	}

	if err := conf.Parse(os.Args[1:], "PPIO", &cfg); err != nil {
		if err == conf.ErrHelpWanted {
			usage, err := conf.Usage("PPIO", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}
	l := log.New(os.Stdout, "ppio :", log.LstdFlags)

	dbCfg := storage.Config{
		User:       cfg.DB.User,
		Password:   cfg.DB.Password,
		Host:       cfg.DB.Host,
		Name:       cfg.DB.Name,
		DisableTLS: cfg.DB.DisableTLS,
	}

	db, dbTidy, err := storage.SetupDB(&dbCfg)
	if err != nil {
		return errors.Wrap(err, "setup database")
	}
	defer dbTidy()
	l.Println("database init OK")

	ch := handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:4200"}))
	srv := ppioHTTP.NewServer(db, l)

	http.ListenAndServe(":9001", ch(srv))
	return nil
}
