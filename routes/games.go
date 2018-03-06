package routes

/*
func getGameHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		ctx := context.Background()
		gameID := vars["gameID"]
		gameGet, err := client.
			Get().
			Index("games").
			Type("doc").
			Id(gameID).
			Do(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		gameJSON, err := gameGet.Source.MarshalJSON()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(gameJSON)
	}

	return http.HandlerFunc(fn)
}

func getGamesHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var games []models.Game
		ctx := context.Background()

		// Fetching count
		count, err := client.Count().
			Index("games").
			Type("doc").
			Do(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Perform search
		docMatchQuery := elastic.NewTermQuery("_type", "doc")
		result, err := client.Search().
			Index("games").
			Query(docMatchQuery).
			From(0).Size(int(count)).
			Do(ctx)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var game models.Game
		for _, item := range result.Each(reflect.TypeOf(game)) {
			if t, ok := item.(models.Game); ok {
				game = models.Game{
					DateTime: t.DateTime,
					Player1:  t.Player1,
					Player2:  t.Player2,
					Score1:   t.Score1,
					Score2:   t.Score2,
				}
				games = append(games, game)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(games)
	}

	return http.HandlerFunc(fn)
}

func addGameHandler(dbConn *sql.DB) http.HandlerFunc {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var game models.Game
		ctx := context.Background()
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

		_, err = client.Index().
			Index("games").
			Type("doc").
			BodyJson(game).
			Refresh("true").
			Do(ctx)

		if err != nil {
			log.Println("Could not insert game")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}

	return http.HandlerFunc(fn)
}
*/
