package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/RaphaelParment/ppio-api/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	Logger   *log.Logger
	Router   *mux.Router
	Database *sql.DB
	Players  *handlers.Players
}

type RequestHandlerFunc func(db *sql.DB, l *log.Logger, w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.Database, a.Logger, w, r)
	}
}

// CreateRouter build the main router
func (a *App) CreateRouter() {
	a.Router = mux.NewRouter()
	s := a.Router.PathPrefix("/ppio").Subrouter()

	getRtr := s.Methods(http.MethodGet).Subrouter()

	getRtr.HandleFunc("/players", a.handleRequest(a.Players.GetPlayers))
	getRtr.HandleFunc("/players/{id:[0-9]+}", a.handleRequest(a.Players.GetPlayer))
	http.Handle("/", a.Router)

}
