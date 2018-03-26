package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"ppio/models"
	"strconv"

	"github.com/gorilla/mux"
)

func getPlayerHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		if playerID, ok := vars["playerID"]; ok {
			var player models.Player
			var id int
			var err error
			if id, err = strconv.Atoi(playerID); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			player.ID = int64(id)
			player.FirstName = vars["first_name"]
			err = player.GetByID(dbConn)

			if err != nil {
				log.Fatalf("Could not get player %v, err: %v", player, err)
			}

			playerJSON, err := json.Marshal(player)

			if err != nil {
				log.Fatalf("Could not parse player: %v, err: %v",
					player, err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(playerJSON)
		} else {
			http.Error(w, "No player ID given", http.StatusBadRequest)
		}
	}

	return http.HandlerFunc(fn)
}

func addPlayerHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var player models.Player
		if _, err := player.Insert(dbConn); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}

	return http.HandlerFunc(fn)
}

func getAllPlayersHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var player models.Player
		players, err := player.GetAll(dbConn)

		if err != nil {
			log.Printf("Could not get players. Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		playerJSON, err := json.Marshal(players)

		if err != nil {
			log.Printf("Could not marshal players: %v, err: %v",
				players, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(playerJSON)
	}

	return http.HandlerFunc(fn)
}
