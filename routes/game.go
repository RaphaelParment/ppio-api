package routes

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"ppio/models"
	"strconv"

	"github.com/gorilla/mux"
)

// GamesResponse Type used to return the amount of rows and the items (limited by default)
type GamesResponse struct {
	Count int64         `json:"count"`
	Items []models.Game `json:"items"`
}

func getGameHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var game models.Game
		vars := mux.Vars(req)
		gameID, err := strconv.Atoi(vars["gameID"])

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		game.ID = int64(gameID)
		err = game.GetByID(dbConn)

		if err != nil {
			log.Fatalf("Could not get game %v, err: %v", game, err)
		}

		gameJSON, err := json.Marshal(game)

		if err != nil {
			log.Fatalf("Could not parse game: %v, err: %v", game, err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(gameJSON)
	}

	return http.HandlerFunc(fn)
}

func getGamesHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var game models.Game
		games, countRow, err := game.GetAll(dbConn)
		if err != nil {
			log.Printf("Could not get games. Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		gameResponse := &GamesResponse{
			Count: countRow,
			Items: games,
		}

		gamesJSON, err := json.Marshal(gameResponse)

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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(reqBody, &game)

		// Making sure request body is well formatted
		if err != nil {
			log.Println("Request body does not match game structure")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = game.Insert(dbConn)

		if err != nil {
			log.Println("Could not insert game.")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}

	return http.HandlerFunc(fn)
}

func updateGameHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var game models.Game
		reqBody, err := ioutil.ReadAll(req.Body)

		vars := mux.Vars(req)
		gameID, err := strconv.ParseInt(vars["gameID"], 10, 64)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = json.Unmarshal(reqBody, &game)

		// TODO check why this fails.
		if gameID != game.ID {
			log.Printf("Supplied game.ID: %d in request body does not match query param gameID: %d",
				game.ID, gameID)
			http.Error(w, "Supplied game.ID in request body does not match query param gameID",
				http.StatusInternalServerError)
			return
		}

		if err != nil {
			log.Printf("Could not unmarshal game from body: %s, err: %v",
				string(reqBody), err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = game.Update(dbConn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write(reqBody)
		}
	}
	return http.HandlerFunc(fn)
}

func deleteGameHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var game models.Game
		vars := mux.Vars(req)
		gameID, err := strconv.ParseInt(vars["gameID"], 10, 64)

		if err != nil {
			log.Printf("Supplied ID is not numberocal")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		game.ID = gameID
		game.GetByID(dbConn)

		gameJSON, err := json.Marshal(game)

		_, err = game.Delete(dbConn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(gameJSON)
	}

	return http.HandlerFunc(fn)
}
