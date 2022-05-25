package main

import (
	"encoding/json"
	"fmt"
	"github.com/RaphaelParment/ppio-api/internal/application/pp_service"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/config"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence/postgres"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence/postgres/entity"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	// "github.com/gorilla/handlers"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	//if err := conf.Parse(os.Args[1:], "PPIO", &cfg); err != nil {
	//	if err == conf.ErrHelpWanted {
	//		usage, err := conf.Usage("PPIO", &cfg)
	//		if err != nil {
	//			return errors.Wrap(err, "generating config usage")
	//		}
	//		fmt.Println(usage)
	//		return nil
	//	}
	//	return errors.Wrap(err, "parsing config")
	//}
	//l := log.New(os.Stdout, "ppio :", log.LstdFlags)

	dbCfg := persistence.Config{
		User:       cfg.DB.User,
		Password:   cfg.DB.Password,
		Host:       cfg.DB.Host,
		Name:       cfg.DB.Name,
		DisableTLS: cfg.DB.DisableTLS,
	}

	db, dbTidy, err := persistence.SetupDB(&dbCfg)
	if err != nil {
		return errors.Wrap(err, "setup database")
	}
	defer dbTidy()

	//ch := handlers.CORS(
	//	handlers.AllowedOrigins([]string{"*"}),
	//	handlers.AllowedMethods([]string{"OPTIONS", "GET", "POST", "PUT", "DELETE"}),
	//	handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}))
	//srv := ppioHTTP.NewServer(db, l)

	matchStore := postgres.NewMatchStore(db)
	matchService := pp_service.NewMatchService(matchStore)

	mux := http.NewServeMux()
	mux.HandleFunc("/match", func(w http.ResponseWriter, r *http.Request) {
		match, err := matchService.HandleFindMatch(r.Context(), matchModel.Id(1))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		matchJSON := entity.MatchToJSON(match)
		m, err := json.Marshal(matchJSON)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		_, err = fmt.Fprintf(w, string(m))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	})

	http.ListenAndServe(":9001", mux)
	return nil
}
