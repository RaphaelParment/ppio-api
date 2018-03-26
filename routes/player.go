package routes

import (
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"fmt"
	"context"
	"database/sql"
	"ppio/models"
	"encoding/json"
)

func GetPlayer(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var player models.Player
		vars := mux.Vars(req)
		gameID := vars["gameID"]

		err := dbConn.QueryRow("SELECT id, player1_id, player2_id, score1," +
			" score2, datetime FROM game WHERE id = $1",
			gameID).Scan(&game.ID, &game.Player1ID, &game.Player2ID,
			&game.Score1, &game.Score2, &game.DateTime)

		if err != nil {
			log.Fatalf("Could not get game %v, err: %v", game, err)
		}


		gameJSON, err := json.Marshal(game)

		if err != nil {
			log.Fatalf("Could not parse game: %v, err: %v",
				game, err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(gameJSON)
	}

	return http.HandlerFunc(fn)
}

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

