package rest

import (
	"encoding/json"
	"fmt"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence/postgres/entity"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//
//import (
//	"database/sql"
//	"github.com/RaphaelParment/ppio-api/internal/domain/player"
//	"github.com/RaphaelParment/ppio-api/internal/domain/player/model"
//	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence"
//	"net/http"
//	"strconv"
//
//	"github.com/gorilla/mux"
//)
//
//// Players struct
//// type Players struct{}
//
//// swagger:route GET /players{id} players getPlayers
//// Returns a list of players
//// responses:
//// 	200: playersResponse
//
//// handlePlayersGet returns all players
//func (s *server) handlePlayersGet() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		s.Logger.Println("GET players")
//		players, err := persistence.GetPlayers(s.DB)
//		if err != nil {
//			s.Logger.Printf("could not get players; %v", err)
//			s.respond(w, r, nil, http.StatusInternalServerError)
//			return
//		}
//		s.respond(w, r, players, http.StatusOK)
//	}
//}
//
//// swagger:route GET /players/{id} players getPlayer
//// Return a single player
//// responses:
////	200: playersResponse
////	404: errorResponse
//
//// handlePlayerGet handler returns a single player based on the <id> from the request
//func (s *server) handlePlayerGet() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		vars := mux.Vars(r)
//		id, err := strconv.Atoi(vars["id"])
//		if err != nil {
//			s.Logger.Printf("failed to convert id: %d to int; %v", id, err)
//			s.respond(w, r, "Wrong id", http.StatusInternalServerError)
//			return
//		}
//		s.Logger.Println("GET player", id)
//		player, err := persistence.GetPlayer(s.DB, id)
//		if err == sql.ErrNoRows {
//			s.Logger.Printf("No player found for id: %d.", id)
//			s.respond(w, r, "Not found", http.StatusNotFound)
//			return
//		}
//		if err != nil {
//			s.Logger.Printf("Unable to get player; %v", err)
//			s.respond(w, r, "Error", http.StatusInternalServerError)
//			return
//		}
//		s.respond(w, r, player, http.StatusOK)
//	}
//}
//
//// swagger:route POST /players players createPlayer
//// Create a new player
////
//// responses:
////	200: playerResponse
////  422: errorValidation
////  501: errorResponse
//
//// handlePlayerAdd handler adds the player from the body of the request to the database
//func (s *server) handlePlayerAdd() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		s.Logger.Print("POST player")
//		player := r.Context().Value(player.KeyPlayer{}).(*model.Player)
//		err := persistence.AddPlayer(s.DB, player)
//		if err != nil {
//			s.Logger.Printf("could not add player; %v", err)
//			s.respond(w, r, "Could not add player", http.StatusInternalServerError)
//			return
//		}
//		s.respond(w, r, player, http.StatusOK)
//	}
//}
//
//// swagger:route PUT /players/{id} players updatePlayer
//// Updates a player
////
//// responses:
////	201: noContentResponse
////  404: errorResponse
////  422: errorValidation
//
//// handlePlayerUpdate handler updates the player with id <id> from the request,
//// by the player from the the request
//func (s *server) handlePlayerUpdate() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		vars := mux.Vars(r)
//		id, err := strconv.Atoi(vars["id"])
//		if err != nil {
//			s.Logger.Printf("Failed to convert id to int; %v", err)
//			s.respond(w, r, "could not update player", http.StatusBadRequest)
//			return
//		}
//		s.Logger.Printf("PUT player id: %d", id)
//		player := r.Context().Value(player.KeyPlayer{}).(*model.Player)
//		success, err := persistence.UpdatePlayer(s.DB, id, player)
//		if err != nil {
//			s.Logger.Printf("could not update player id: %d %v", id, err)
//			s.respond(w, r, "Could not update", http.StatusInternalServerError)
//			return
//		}
//		if !success {
//			s.Logger.Printf("no player to update for id: %d", id)
//			s.respond(w, r, "Player not found", http.StatusNotFound)
//			return
//		}
//		s.respond(w, r, player, http.StatusOK)
//	}
//}
//
//// swagger:route DELETE /players/{id} players deletePlayer
//// Delete a player
////
//// responses:
////	201: noContentResponse
////  404: errorResponse
////  501: errorResponse
//
//// handlePlayerDelete handler deletes the player with id <id> from the request
//func (s *server) handlePlayerDelete() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		vars := mux.Vars(r)
//		id, err := strconv.Atoi(vars["id"])
//		if err != nil {
//			s.Logger.Printf("Failed to convert id to int; %v", err)
//			s.respond(w, r, "Failed to delete player", http.StatusBadRequest)
//			return
//		}
//		s.Logger.Printf("DELETE player id: %d", id)
//		success, err := persistence.DeletePlayer(s.DB, id)
//		if err != nil {
//			s.Logger.Printf("Failed to delete player id: %d; %v", id, err)
//			s.respond(w, r, "Failed to delete player", http.StatusInternalServerError)
//			return
//		}
//		if !success {
//			s.Logger.Printf("No player to delte with id: %d", id)
//			s.respond(w, r, "Player not found", http.StatusNotFound)
//			return
//		}
//		s.respond(w, r, nil, http.StatusOK)
//	}
//
//}

func (s *server) HandleFindPlayers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestId, found := mux.Vars(r)["id"]
		if !found {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(requestId)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		match, err := s.matchService.HandleFindMatch(r.Context(), matchModel.Id(id))
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
	}
}
