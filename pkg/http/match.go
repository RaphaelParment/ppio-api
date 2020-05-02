package http

import (
	"net/http"

	"github.com/RaphaelParment/ppio-api/pkg/storage"
)

func (s *server) handletMatchesGet() http.HandlerFunc {
	statusCode := http.StatusOK
	matches, err := storage.GetMatches(s.DB)
	if err != nil {
		statusCode = http.StatusInternalServerError
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, matches, statusCode)
	}
}
