package routes

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/RaphaelParment/ppio-api/models"

	"github.com/gorilla/mux"
)

// CountResponse Type used to return the amount of rows and the items (limited by default)
type CountResponse struct {
	Count int64           `json:"count"`
	Items []models.Player `json:"items"`
}

func parsePlayerParameters(vars url.Values) (map[string]interface{}, error) {
	filters := make(map[string]interface{})

	playerFirstName := vars.Get("firstName")
	if playerFirstName != "" {
		filters["firstName"] = playerFirstName
	}
	playerLastName := vars.Get("lastName")
	if playerLastName != "" {
		filters["lastName"] = playerLastName
	}

	points := vars.Get("points")
	if points != "" {
		pointsNr, err := strconv.Atoi(points)
		if err != nil {
			log.Printf("Could not parse the value of the points. Err : %v", err)
		} else {
			filters["points"] = pointsNr
		}
	}

	return filters, nil
}

func getPlayerHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		if playerID, ok := vars["playerID"]; ok {
			var player models.Player
			var err error

			player.ID = playerID
			err = player.GetByID(dbConn)

			if err != nil {
				log.Printf("Could not get player %v, err: %v\n", player, err)
				http.Error(w, "Player ID not found.\n",
					http.StatusNotFound)
				return
			}
			playerJSON, err := json.Marshal(player)

			if err != nil {
				log.Printf("Could not parse player: %v, err: %v",
					player, err)
				http.Error(w, "Could not parse player.\n",
					http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(playerJSON)

		} else {
			http.Error(w, "No player ID given", http.StatusBadRequest)
		}
	}

	return http.HandlerFunc(fn)
}

func getPlayerGamesHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		urlVars := req.URL.Query()
		filters := make(map[string]interface{})
		if playerID, ok := vars["playerID"]; ok {
			var game models.Game
			var err error
			filters["fromPlayerID"] = playerID

			limit, offset := parseLimitAndOffset(urlVars)
			filters["limit"] = limit
			filters["offset"] = offset

			games, countRow, err := game.GetAll(dbConn, filters)

			gamesResponse := &GamesResponse{
				Count: countRow,
				Items: games,
			}

			gamesJSON, err := json.Marshal(gamesResponse)

			if err != nil {
				log.Fatalf("Could not parse player: %v, err: %v",
					gamesResponse, err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(gamesJSON)
		} else {
			http.Error(w, "No player ID given", http.StatusBadRequest)
		}
	}

	return http.HandlerFunc(fn)
}

func addPlayerHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var player models.Player

		reqBody, err := ioutil.ReadAll(req.Body)

		if err != nil {
			log.Print("Could not read request body while adding player\n")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(reqBody, &player)

		// Making sure request body is well formatted
		if err != nil {
			log.Print("Request body does not match player structure\n")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := player.Insert(dbConn); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		playerJSON, err := json.Marshal(player)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(playerJSON)
	}

	return http.HandlerFunc(fn)
}

func getAllPlayersHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		vars := req.URL.Query()
		filters, err := parsePlayerParameters(vars)
		if err != nil {
			log.Printf("Could not parse the game query parameters")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		limit, offset := parseLimitAndOffset(vars)
		filters["limit"] = limit
		filters["offset"] = offset

		var player models.Player
		players, countRow, err := player.GetAll(dbConn, filters)
		if err != nil {
			log.Printf("Could not get players. Error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		playersResponse := &CountResponse{
			Count: countRow,
			Items: players,
		}
		playerJSON, err := json.Marshal(playersResponse)

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

func updatePlayerHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		if playerID, ok := vars["playerID"]; ok {
			var player models.Player
			var err error

			player.ID = playerID
			requestBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Printf("Could not handle PUT request for update player. Error: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = json.Unmarshal(requestBytes, &player)
			if err != nil {
				log.Printf("Could not unmarshal players from body: %s, err: %v",
					string(requestBytes), err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			_, err = player.Update(dbConn)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)

			} else {
				w.Header().Set("Content-Type", "application/json")
				w.Write(requestBytes)
			}
		} else {
			http.Error(w, "No player ID given", http.StatusBadRequest)
		}
	}

	return http.HandlerFunc(fn)
}
