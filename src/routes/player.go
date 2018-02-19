package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	elastic "gopkg.in/olivere/elastic.v5"
)

func getPlayerHandler(client *elastic.Client) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ctx := context.Background()
		playerID := vars["playerID"]
		playerGet, err := client.
			Get().
			Index("players").
			Type("doc").
			Id(playerID).
			Do(ctx)
		if err != nil {
			log.Printf("Could not get the player by ID: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if playerGet.Found {
			fmt.Printf("Got the document: %v", playerGet.Source)
		}

		playerJSON, err := playerGet.Source.MarshalJSON()
		if err != nil {
			log.Printf("Could not marshal the retrieved user: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(playerJSON)

	}
	return http.HandlerFunc(fn)
}
