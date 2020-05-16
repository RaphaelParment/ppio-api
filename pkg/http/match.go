package http

import (
	"net/http"

	"github.com/RaphaelParment/ppio-api/pkg/core"
	"github.com/RaphaelParment/ppio-api/pkg/storage"
)

func (s *server) handleMatchesGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Println("GET matches")
		matches, err := storage.GetMatches(s.DB)
		if err != nil {
			s.Logger.Printf("could not get matches; %v", err)
			s.respond(w, r, nil, http.StatusInternalServerError)
			return
		}
		s.respond(w, r, matches, http.StatusOK)
	}
}

func (s *server) handleMatchAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Logger.Print("POST match")
		match := r.Context().Value(core.KeyMatch{}).(*core.Match)
		err := storage.AddMatch(s.DB, match)
		if err != nil {
			s.Logger.Printf("could not add match; %v", err)
			s.respond(w, r, "Could not add match", http.StatusInternalServerError)
			return
		}
		s.respond(w, r, match, http.StatusOK)
	}
}
