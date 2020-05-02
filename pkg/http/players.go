// Package classification: PPIO API
//
// Documentation for Player API
//
// 	Schemes: http
// 	BasePath: /ppio
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
// swagger:meta
package http

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/RaphaelParment/ppio-api/pkg/core"
	"github.com/RaphaelParment/ppio-api/pkg/storage"
	"github.com/gorilla/mux"
)

// Players struct
// type Players struct{}

// KeyPlayer struct is used as key for the request context
type KeyPlayer struct{}

// swagger:route GET /players{id} players getPlayers
// Returns a list of players
// responses:
// 	200: playersResponse

// handlePlayersGet returns all players
func (s *server) handlePlayersGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Println("GET players")
		players, err := storage.GetPlayers(s.DB)
		if err != nil {
			s.Logger.Printf("could not get players; %v", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		s.respond(w, r, players, http.StatusOK)
	}
}

// swagger:route GET /players/{id} players getPlayer
// Return a single player
// responses:
//	200: playersResponse
//	404: errorResponse

// handlePlayerGet handler returns a single player based on the <id> from the request
func (s *server) handlePlayerGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.Logger.Printf("failed to convert id: %d to int; %v", id, err)
			s.respond(w, r, "Wrong id", http.StatusInternalServerError)
			return
		}
		s.Logger.Println("GET player", id)
		player, err := storage.GetPlayer(s.DB, id)
		if err == sql.ErrNoRows {
			s.Logger.Printf("No player found for id: %d.", id)
			s.respond(w, r, "Not found", http.StatusNotFound)
			return
		}
		if err != nil {
			s.Logger.Printf("Unable to get player; %v", err)
			s.respond(w, r, "Error", http.StatusInternalServerError)
			return
		}
		s.respond(w, r, player, http.StatusOK)
	}
}

// swagger:route POST /players players createPlayer
// Create a new player
//
// responses:
//	200: playerResponse
//  422: errorValidation
//  501: errorResponse

// handlePlayerAdd handler adds the player from the body of the request to the database
func (s *server) handlePlayerAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Print("POST player")
		player := r.Context().Value(KeyPlayer{}).(core.Player)
		err := storage.AddPlayer(s.DB, &player)
		if err != nil {
			s.Logger.Printf("could not add player; %v", err)
			s.respond(w, r, "Could not add player", http.StatusInternalServerError)
			return
		}
		s.respond(w, r, player, http.StatusOK)
	}
}

// swagger:route PUT /players/{id} players updatePlayer
// Updates a player
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// handlePlayerUpdate handler updates the player with id <id> from the request,
// by the player from the the request
func (s *server) handlePlayerUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.Logger.Printf("Failed to convert id to int; %v", err)
			s.respond(w, r, "could not update player", http.StatusBadRequest)
			return
		}
		s.Logger.Printf("PUT player id: %d", id)
		player := r.Context().Value(KeyPlayer{}).(core.Player)
		success, err := storage.UpdatePlayer(s.DB, id, &player)
		if err != nil {
			s.Logger.Printf("could not update player id: %d %v", id, err)
			s.respond(w, r, "Could not update", http.StatusInternalServerError)
			return
		}
		if !success {
			s.Logger.Printf("no player to update for id: %d", id)
			s.respond(w, r, "Player not found", http.StatusNotFound)
			return
		}
		s.respond(w, r, player, http.StatusOK)
	}
}

// swagger:route DELETE /players/{id} players deletePlayer
// Delete a player
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// handlePlayerDelete handler deletes the player with id <id> from the request
func (s *server) handlePlayerDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.Logger.Printf("Failed to convert id to int; %v", err)
			s.respond(w, r, "Failed to delete player", http.StatusBadRequest)
			return
		}
		s.Logger.Printf("DELETE player id: %d", id)
		success, err := storage.DeletePlayer(s.DB, id)
		if err != nil {
			s.Logger.Printf("Failed to delete player id: %d; %v", id, err)
			s.respond(w, r, "Failed to delete player", http.StatusInternalServerError)
			return
		}
		if !success {
			s.Logger.Printf("No player to delte with id: %d", id)
			s.respond(w, r, "Player not found", http.StatusNotFound)
			return
		}
		s.respond(w, r, nil, http.StatusOK)
	}

}

// playerValid middleware checks if the player specified in the
// body of the request is correct.
func (s *server) playerValid(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		player := core.Player{}
		err := player.FromJSON(r.Body)
		if err != nil {
			s.Logger.Printf("could not convert to player; %v", err)
			s.respond(w, r, "Wrong JSON", http.StatusBadRequest)
			return
		}
		err = player.Validate()
		if err != nil {
			s.Logger.Printf("could not validate player; %v", err)
			s.respond(w, r, "Error reading player", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyPlayer{}, player)
		r = r.WithContext(ctx)

		h(w, r)
	}
}
