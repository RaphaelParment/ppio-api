package rest

//
//import (
//	"github.com/RaphaelParment/ppio-api/internal/domain/match"
//	"github.com/RaphaelParment/ppio-api/internal/domain/match/model"
//	model2 "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
//	"net/http"
//
//	"github.com/gorilla/mux"
//)
//
//func (s *server) routes() {
//
//	getRtr := s.Router.Methods(http.MethodGet).Subrouter()
//	postRtr := s.Router.Methods(http.MethodPost).Subrouter()
//	optsRtr := s.Router.Methods(http.MethodOptions).Subrouter()
//	putRtr := s.Router.Methods(http.MethodPut).Subrouter()
//	delRtr := s.Router.Methods(http.MethodDelete).Subrouter()
//
//	getRtr.HandleFunc("/players", s.handlePlayersGet())
//	getRtr.HandleFunc("/players/{id:[0-9]+}", s.handlePlayerGet())
//	getRtr.HandleFunc("/matches", s.handleMatchesGet())
//	getRtr.HandleFunc("/results/{id:[0-9]+}", s.handleMatchResultGet())
//	getRtr.HandleFunc("/scores/{id:[0-9]+}", s.handleMatchGamesScoresGet())
//
//	s.Router.Handle("/swagger.yaml", s.handleRawDocsGet()).Methods(http.MethodGet)
//	s.Router.Handle("/docs", s.handleDocsGet()).Methods(http.MethodGet)
//
//	optsRtr.HandleFunc("/players", s.handlePreflight())
//	optsRtr.HandleFunc("/matches", s.handlePreflight())
//	optsRtr.HandleFunc("/results", s.handlePreflight())
//	optsRtr.HandleFunc("/scores", s.handlePreflight())
//
//	postRtr.HandleFunc("/players", s.resourceValid(s.handlePlayerAdd(), &model2.Player{}))
//	postRtr.HandleFunc("/matches", s.resourceValid(s.handleMatchAdd(), &model.Match{}))
//	postRtr.HandleFunc("/results", s.resourceValid(s.handleMatchResultAdd(), &match.MatchResult{}))
//	postRtr.HandleFunc("/scores", s.resourceValid(s.handleMatchGamesScoresAdd(), &match.GameScores{}))
//
//	putRtr.HandleFunc("/players/{id:[0-9]+}", s.resourceValid(s.handlePlayerUpdate(), &model2.Player{}))
//
//	delRtr.HandleFunc("/players/{id:[0-9]+}", s.handlePlayerDelete())
//
//	s.Router.Use(mux.CORSMethodMiddleware(s.Router))
//}
