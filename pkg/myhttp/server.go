package myhttp

import (
	"database/sql"
	"encoding/json"
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	DB     *sql.DB
	Logger *log.Logger
	Router *mux.Router
}

func NewServer(db *sql.DB, l *log.Logger) *server {
	srv := server{
		DB:     db,
		Logger: l,
		Router: mux.NewRouter(),
	}
	srv.routes()
	return &srv
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, "failed to convert to JSON", http.StatusInternalServerError)
			return
		}
		b, _ := data.([]byte)
		w.Write(b)
	}
}

func (s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
