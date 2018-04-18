package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

// GetRouter Instantiate the router and mounts all the handler for the Zone30  routes.
func GetRouter(dbConn *sql.DB) *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/ppio").Subrouter()
	s.HandleFunc("/players/{playerID}", getPlayerHandler(dbConn)).Methods(http.MethodGet)
	s.HandleFunc("/players/{playerID}", updatePlayerHandler(dbConn)).Methods(http.MethodPut)
	s.HandleFunc("/players", getAllPlayersHandler(dbConn)).Methods(http.MethodGet)
	s.HandleFunc("/players", addPlayerHandler(dbConn)).Methods(http.MethodPost)
	s.HandleFunc("/games/{gameID}", getGameHandler(dbConn)).Methods(http.MethodGet)
	s.HandleFunc("/games/{gameID}", updateGameHandler(dbConn)).Methods(http.MethodPut)
	s.HandleFunc("/games/{gameID}", deleteGameHandler(dbConn)).Methods(http.MethodDelete)
	s.HandleFunc("/games", getGamesHandler(dbConn)).Methods(http.MethodGet)
	s.HandleFunc("/games", addGameHandler(dbConn)).Methods(http.MethodPost)
	http.Handle("/", r)
	return r
}