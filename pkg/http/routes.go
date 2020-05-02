package http

import (
	"net/http"
)

func (s *server) routes() {

	getRtr := s.Router.Methods(http.MethodGet).Subrouter()
	postRtr := s.Router.Methods(http.MethodPost).Subrouter()
	putRtr := s.Router.Methods(http.MethodPut).Subrouter()
	delRtr := s.Router.Methods(http.MethodDelete).Subrouter()

	getRtr.HandleFunc("/players", s.handlePlayersGet())
	getRtr.HandleFunc("/players/{id:[0-9]+}", s.handlePlayerGet())
	getRtr.Handle("/swagger.yaml", s.handleRawDocsGet())
	getRtr.Handle("/docs", s.handleDocsGet())

	postRtr.HandleFunc("/players", s.playerValid(s.handlePlayerAdd()))

	putRtr.HandleFunc("/players/{id:[0-9]+}", s.playerValid(s.handlePlayerUpdate()))

	delRtr.HandleFunc("/players/{id:[0-9]+}", s.handlePlayerDelete())

}
