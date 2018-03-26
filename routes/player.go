package routes

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"ppio/models"

	"github.com/gorilla/mux"
)

func getPlayerHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var player models.Player
		vars := mux.Vars(req)
		if playerID, ok := vars["playerID"]; ok {
			err := dbConn.QueryRow("SELECT id, first_name, last_name, points FROM player WHERE id = $1",
				playerID).Scan(&player.ID, &player.FirstName, &player.LastName, &player.Points)

			if err != nil {
				log.Fatalf("Could not get game %v, err: %v", player, err)
			}

			playerJSON, err := json.Marshal(player)

			if err != nil {
				log.Fatalf("Could not parse game: %v, err: %v",
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
