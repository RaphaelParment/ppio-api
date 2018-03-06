package main

import (
	"database/sql"
	"log"
	"os"
	"ppio/utils"

	"fmt"

	_ "github.com/lib/pq"
)

/**
Function which inserts dummy data into the database.
*/
func initialiseDb() *sql.DB {

	db, err := sql.Open("postgres", "postgresql://ppio_user@localhost:26257/ppio?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	players := utils.GetPlayers()

	for _, player := range players {
		lastID := player.Insert(db)
		player.ID = lastID
	}
	fmt.Println("Players inserted")

	games := utils.GenerateGames(players)

	for _, game := range games {

		_ = game.Insert(db)
	}
	fmt.Println("Games inserted")

	return db
}

func main() {

	end := make(chan bool)

	dbConn := initialiseDb()
	defer dbConn.Close()

	// Handle the routes with gorillamux
	/*
		go func() {
			http.ListenAndServe(":9000", routes.GetRouter(client))
		}()
	*/

	// TODO check where it could be sent ?
	// Handle the termination of the program properly.
	for {
		select {
		case StopProg := <-end:
			if StopProg {
				os.Exit(0)
			}
		}
	}

}
