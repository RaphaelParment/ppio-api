package http

import (
	"net/http"

	"github.com/RaphaelParment/ppio-api/pkg/core"
)

func (s *server) routes() {

	getRtr := s.Router.Methods(http.MethodGet).Subrouter()
	postRtr := s.Router.Methods(http.MethodPost).Subrouter()
	putRtr := s.Router.Methods(http.MethodPut).Subrouter()
	delRtr := s.Router.Methods(http.MethodDelete).Subrouter()

	getRtr.HandleFunc("/players", s.handlePlayersGet())
	getRtr.HandleFunc("/players/{id:[0-9]+}", s.handlePlayerGet())
	getRtr.HandleFunc("/matches", s.handleMatchesGet())
	getRtr.HandleFunc("/results/{id:[0-9]+}", s.handleMatchResultGet())
	getRtr.HandleFunc("/scores/{id:[0-9]+}", s.handleMatchGamesScoresGet())

	getRtr.Handle("/swagger.yaml", s.handleRawDocsGet())
	getRtr.Handle("/docs", s.handleDocsGet())

	postRtr.HandleFunc("/players", s.resourceValid(s.handlePlayerAdd(), &core.Player{}))
	postRtr.HandleFunc("/matches", s.resourceValid(s.handleMatchAdd(), &core.Match{}))
	postRtr.HandleFunc("/results", s.resourceValid(s.handleMatchResultAdd(), &core.MatchResult{}))
	postRtr.HandleFunc("/scores", s.resourceValid(s.handleMatchGamesScoresAdd(), &core.GameScores{}))

	putRtr.HandleFunc("/players/{id:[0-9]+}", s.resourceValid(s.handlePlayerUpdate(), &core.Player{}))

	delRtr.HandleFunc("/players/{id:[0-9]+}", s.handlePlayerDelete())

}
