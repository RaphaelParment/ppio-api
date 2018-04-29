package routes

import (
	"database/sql"
	"log"
	"net/http"
	"ppio/models"
	"strconv"

	"github.com/gorilla/mux"
)

func parseLimitAndOffset(vars map[string]string) (int, int) {
	limit, ok := vars["limit"]
	var limitNr, offsetNr int
	var err error
	if ok {
		limitNr, err = strconv.Atoi(limit)
		if err != nil {
			log.Printf("Could not parse the value of limit paramter. Err: %v", err)
			limitNr = models.DefaultLimit
		} else {
			limitNr = models.DefaultLimit
		}

		offset, ok := vars["offset"]
		if ok {
			offsetNr, err = strconv.Atoi(offset)
			if err != nil {
				log.Printf("Could not parse the value of offset parameter. Err: %v", err)
				offsetNr = models.DefaultOffset
			}
		} else {
			offsetNr = models.DefaultOffset
		}
	}
	return limitNr, offsetNr
}

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
