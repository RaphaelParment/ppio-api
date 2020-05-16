package http

import (
	"net/http"
	"strconv"

	"github.com/RaphaelParment/ppio-api/pkg/core"
	"github.com/RaphaelParment/ppio-api/pkg/storage"
	"github.com/gorilla/mux"
)

func (s *server) handleMatchResultGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.Logger.Printf("failed to convert id: %d to int; %v", id, err)
			s.respond(w, r, "Wrong id", http.StatusInternalServerError)
			return
		}
		s.Logger.Printf("GET match %d result", id)
		result, err := storage.GetMatchResult(s.DB, id)
		if err != nil {
			s.Logger.Printf("could not get result for match %d; %v", id, err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		s.respond(w, r, result, http.StatusOK)
	}
}

func (s *server) handleMatchResultAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Print("POST match result")
		result := r.Context().Value(core.KeyMatchResult{}).(*core.MatchResult)
		err := storage.AddMatchResult(s.DB, result)
		if err != nil {
			s.Logger.Printf("could not add match result; %v", err)
			s.respond(w, r, "Could not add match result", http.StatusInternalServerError)
			return
		}
		s.respond(w, r, result, http.StatusOK)
	}
}
