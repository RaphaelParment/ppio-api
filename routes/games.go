package routes

import (
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	"ppio/models"
	"log"
	"encoding/json"

	"io/ioutil"
)

func getGameHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var game models.Game
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

func getGamesHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var games []models.Game
		var game models.Game

		results, err := dbConn.Query("SELECT id, player1_id, player2_id, score1," +
			" score2, datetime FROM game")

		if err != nil {
			log.Fatalf("Could not get all games, err: %v", err)
		}

		for results.Next() {
			err := results.Scan(&game.ID, &game.Player1ID, &game.Player2ID,
				&game.Score1, &game.Score2, &game.DateTime)
			if err != nil {
				log.Fatalf("Could not parse game, err: %v", err)
			}
			games = append(games, game)
		}

		gamesJSON, err := json.Marshal(games)
		if err != nil {
			log.Fatalf("Could not parse games to JSON, err: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(gamesJSON)

	}

	return http.HandlerFunc(fn)
}

func addGameHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var game models.Game
		reqBody, err := ioutil.ReadAll(req.Body)

		if err != nil {
			log.Println("Could not read request body while adding game.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(reqBody, &game)

		// Making sure request body is well formatted
		if err != nil {
			log.Println("Request body does not match game structure")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = dbConn.Exec("INSERT INTO game (player1_id, player2_id, score1, score2, datetime)" +
			"VALUES ($1, $2, $3, $4, NOW())", game.Player1ID, game.Player2ID,
				game.Score1, game.Score2)

		if err != nil {
			log.Println("Could not insert game.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}

	return http.HandlerFunc(fn)
}
