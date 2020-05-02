package myhttp

func (s *server) routes() {
	s.Router.HandleFunc("/players", s.handlePlayersGet())
}
