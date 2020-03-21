package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/RaphaelParment/ppio-api/database"
	"github.com/gorilla/mux"
)

// Players struct
type Players struct {
}

// NewPlayers func
func NewPlayers() *Players {
	return &Players{}
}

// GetPlayers func
func (p *Players) GetPlayers(db *sql.DB, l *log.Logger, rw http.ResponseWriter, r *http.Request) {
	l.Println("handlers : GET players")
	players, err := database.GetPlayers(db)
	if err != nil {
		http.Error(rw, "Unable to fetch players", http.StatusInternalServerError)
		return
	}
	err = players.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert players to json", http.StatusInternalServerError)
		return
	}
}

// GetPlayer handler to fech a single player
func (p *Players) GetPlayer(db *sql.DB, l *log.Logger, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "failed to convert id to int", http.StatusBadRequest)
		return
	}
	l.Println("handlers : GET player", id)
	player, err := database.GetPlayer(db, id)
	if err == sql.ErrNoRows {
		http.Error(rw, "Not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Unable to get player", http.StatusInternalServerError)
		return
	}
	err = player.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert player to json", http.StatusInternalServerError)
		return
	}
}
