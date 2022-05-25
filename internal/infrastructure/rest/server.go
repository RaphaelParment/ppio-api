package rest

//
//import (
//	"context"
//	"database/sql"
//	"encoding/json"
//	"github.com/RaphaelParment/ppio-api/internal/domain"
//	"log"
//
//	"net/http"
//
//	"github.com/gorilla/mux"
//)
//
//type server struct {
//	DB     *sql.DB
//	Logger *log.Logger
//	Router *mux.Router
//}
//
//func NewServer(db *sql.DB, l *log.Logger) *server {
//	srv := server{
//		DB:     db,
//		Logger: l,
//		Router: mux.NewRouter(),
//	}
//	srv.routes()
//	return &srv
//}
//
//func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	s.Router.ServeHTTP(w, r)
//}
//
//func (s *server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
//	w.Header().Add("Content-Type", "application/json")
//	w.WriteHeader(status)
//	if data != nil {
//		err := json.NewEncoder(w).Encode(data)
//		if err != nil {
//			http.Error(w, "failed to convert to JSON", http.StatusInternalServerError)
//			return
//		}
//		b, _ := data.([]byte)
//		w.Write(b)
//	}
//}
//
//func (s *server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
//	return json.NewDecoder(r.Body).Decode(v)
//}
//
//func (s *server) resourceValid(h http.HandlerFunc, res domain.Resource) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		err := res.FromJSON(r.Body)
//		if err != nil {
//			s.Logger.Printf("could not convert to resource %v; %v", res, err)
//			s.respond(w, r, "Wrong JSON", http.StatusBadRequest)
//			return
//		}
//		err = res.Validate()
//		if err != nil {
//			s.Logger.Printf("could not validate resource %v; %v", res, err)
//			s.respond(w, r, "Error reading match", http.StatusBadRequest)
//			return
//		}
//
//		ctx := context.WithValue(r.Context(), res.GetKey(), res)
//		r = r.WithContext(ctx)
//
//		h(w, r)
//	}
//}
//
//func (s *server) handlePreflight() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
//	}
//}
