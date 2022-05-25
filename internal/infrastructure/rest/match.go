package rest

//
//import (
//	"github.com/RaphaelParment/ppio-api/internal/domain"
//	"github.com/RaphaelParment/ppio-api/internal/domain/match/model"
//	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence"
//	"net/http"
//)
//
//func (s *server) handleMatchesGet() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		s.Logger.Println("GET matches")
//		matches, err := persistence.GetMatches(s.DB)
//		if err != nil {
//			s.Logger.Printf("could not get matches; %v", err)
//			s.respond(w, r, nil, http.StatusInternalServerError)
//			return
//		}
//		s.respond(w, r, matches, http.StatusOK)
//	}
//}
//
//func (s *server) handleMatchAdd() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		s.Logger.Print("POST match")
//		match := r.Context().Value(domain.KeyMatch{}).(*model.Match)
//		err := persistence.AddMatch(s.DB, match)
//		if err != nil {
//			s.Logger.Printf("could not add match; %v", err)
//			s.respond(w, r, "Could not add match", http.StatusInternalServerError)
//			return
//		}
//		s.respond(w, r, match, http.StatusOK)
//	}
//}
