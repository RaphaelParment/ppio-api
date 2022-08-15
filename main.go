package main

import (
	"github.com/RaphaelParment/ppio-api/internal/application/pp_service"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/config"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence/postgres"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/rest"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	// "github.com/gorilla/handlers"
	"github.com/pkg/errors"
)

func main() {
	appLogger := log.New(os.Stdout, "", log.LstdFlags)
	if err := run(appLogger); err != nil {
		appLogger.Println(err)
		os.Exit(1)
	}
}

func run(logger *log.Logger) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	dbCfg := persistence.Config{
		User:       cfg.DB.User,
		Password:   cfg.DB.Password,
		Host:       cfg.DB.Host + cfg.DB.Port,
		Name:       cfg.DB.Name,
		DisableTLS: cfg.DB.DisableTLS,
	}

	db, dbTidy, err := persistence.ConnectAndTidy(&dbCfg)
	if err != nil {
		return errors.Wrap(err, "setup database")
	}
	defer dbTidy(logger)

	matchStore := postgres.NewMatchStore(logger, db)
	matchService := pp_service.NewMatchService(matchStore)

	server := rest.NewServer(logger, matchService)

	mux := http.NewServeMux()

	mux.HandleFunc("/matches/{id}", server.HandleOneMatch())
	mux.HandleFunc("/matches", server.HandleMatches())

	http.ListenAndServe(cfg.Http.Port, mux)
	return nil
}
