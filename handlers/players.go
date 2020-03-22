package handlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/RaphaelParment/ppio-api/data"
	"github.com/RaphaelParment/ppio-api/database"
	"github.com/gorilla/mux"
)

// Players struct
type Players struct{}

// KeyPlayer struct is used as key for the request context
type KeyPlayer struct{}

// NewPlayers func
func NewPlayers() *Players {
	return &Players{}
}

// GetPlayers handler returns all players
func (p *Players) GetPlayers(db *sql.DB, l *log.Logger, rw http.ResponseWriter, r *http.Request) {
	l.Println("handlers : GET players")
	players, err := database.GetPlayers(db)
	if err != nil {
		l.Printf("Failed to fetch players. err: %s", err.Error())
		http.Error(rw, "Could not retrieve players", http.StatusInternalServerError)
		return
	}
	err = players.ToJSON(rw)
	if err != nil {
		l.Printf("Failed to convert players to JSON. err: %s", err.Error())
		http.Error(rw, "Could not retrieve players", http.StatusInternalServerError)
		return
	}
}

// GetPlayer handler returns a single player based on the <id> from the request
func (p *Players) GetPlayer(db *sql.DB, l *log.Logger, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		l.Printf("failed to convert id: %d to int. err: %s", id, err.Error())
		http.Error(rw, "Wrong id", http.StatusBadRequest)
		return
	}
	l.Println("handlers : GET player", id)
	player, err := database.GetPlayer(db, id)
	if err == sql.ErrNoRows {
		l.Printf("No player found for id: %d.", id)
		http.Error(rw, "Not found", http.StatusNotFound)
		return
	}
	if err != nil {
		l.Printf("Unable to get player. err: %s", err.Error())
		http.Error(rw, "Error", http.StatusInternalServerError)
		return
	}
	err = player.ToJSON(rw)
	if err != nil {
		l.Printf("Unable to convert player to JSON. err: %s", err.Error())
		http.Error(rw, "Error", http.StatusInternalServerError)
		return
	}
}

// AddPlayer handler adds the player from the body of the request to the database
func (p *Players) AddPlayer(db *sql.DB, l *log.Logger, rw http.ResponseWriter, r *http.Request) {
	l.Print("handlers : POST player")
	player := r.Context().Value(KeyPlayer{}).(data.Player)
	err := database.AddPlayer(db, &player)
	if err != nil {
		l.Printf("Could not add player: %s", err.Error())
		http.Error(rw, "Failed to add player", http.StatusInternalServerError)
		return
	}
	if err := player.ToJSON(rw); err != nil {
		l.Printf("Failed to unmarshal player. err: %s", err.Error())
		http.Error(rw, "Failed to add player", http.StatusInternalServerError)
		return
	}
}

// UpdatePlayer handler updates the player with id <id> from the request,
// by the player from the the request
func (p *Players) UpdatePlayer(db *sql.DB, l *log.Logger, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		l.Printf("Failed to convert id to int. err: %s", err.Error())
		http.Error(rw, "Failed to update player", http.StatusBadRequest)
		return
	}
	l.Printf("handlers : PUT player id: %d", id)
	player := r.Context().Value(KeyPlayer{}).(data.Player)
	success, err := database.UpdatePlayer(db, id, &player)
	if err != nil {
		l.Printf("Failed to update player id: %d. err: %s", id, err.Error())
		http.Error(rw, "Failed to update player", http.StatusInternalServerError)
		return
	}
	if !success {
		l.Printf("No player to update for id: %d", id)
		http.Error(rw, "Player not found", http.StatusNotFound)
		return
	}

	if err := player.ToJSON(rw); err != nil {
		l.Printf("Failed to unmarshal player. err: %s", err.Error())
		http.Error(rw, "Failed to update player", http.StatusInternalServerError)
		return
	}
}

// DeletePlayer handler deletes the player with id <id> from the request
func (p *Players) DeletePlayer(db *sql.DB, l *log.Logger, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		l.Printf("Failed to convert id to int. err: %s", err.Error())
		http.Error(rw, "Failed to delete player", http.StatusBadRequest)
		return
	}
	l.Printf("handlers : DELETE player id: %d", id)
	success, err := database.DeletePlayer(db, id)
	if err != nil {
		l.Printf("Failed to delete player id: %d. err: %s", id, err.Error())
		http.Error(rw, "Failed to delete player", http.StatusInternalServerError)
		return
	}
	if !success {
		l.Printf("No player to delte with id: %d", id)
		http.Error(rw, "Player not found", http.StatusNotFound)
		return
	}
}

// MiddelwarePlayerValidation function does the JSON to player unmarshalling
// and sets the player object onto the context of the request
func (p *Players) MiddelwarePlayerValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		player := data.Player{}

		err := player.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Wrong JSON", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyPlayer{}, player)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
