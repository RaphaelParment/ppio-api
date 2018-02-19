package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	elastic "gopkg.in/olivere/elastic.v5"
)

// GetRouter Instantiate the router and mounts all the handler for the Zone30  routes.
func GetRouter(client *elastic.Client) *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/ppio").Subrouter()
	s.HandleFunc("/players/{playerID}", getPlayerHandler(client)).Methods(http.MethodGet)
	http.Handle("/", r)
	return r
}
