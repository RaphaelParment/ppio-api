package http

import (
	"net/http"
	"strconv"

	"github.com/RaphaelParment/ppio-api/pkg/core"
	"github.com/RaphaelParment/ppio-api/pkg/storage"
	"github.com/gorilla/mux"
)

func (s *server) handleMatchGamesScoresGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.Logger.Printf("failed to convert id: %d to int; %v", id, err)
			s.respond(w, r, "Wrong id", http.StatusInternalServerError)
			return
		}
		scores, err := storage.GetMatchGamesScores(s.DB, id)
		if err != nil {
			s.Logger.Printf("could not get game scores for match %d; %v", id, err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		s.respond(w, r, scores, http.StatusOK)
	}
}

func (s *server) handleMatchGamesScoresAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Print("POST match games scores")
		scores := r.Context().Value(core.KeyMatchGamesScores{}).(*core.GameScores)
		err := storage.AddMatchGamesScores(s.DB, *scores)
		if err != nil {
			s.Logger.Printf("could not add match games scores; %v", err)
			s.respond(w, r, "Could not add match games scores", http.StatusInternalServerError)
			return
		}
		s.respond(w, r, scores, http.StatusOK)
	}
}
