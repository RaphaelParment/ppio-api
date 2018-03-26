package routes


/*
func GetPlayer(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var player models.Player
		vars := mux.Vars(req)
		if playerID, ok := vars["playerID"]; ok {
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
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(gameJSON)
	}

	return http.HandlerFunc(fn)
}
*/
