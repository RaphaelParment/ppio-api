package main

import (
	"database/sql"
	"log"
	"os"
	"ppio/utils"

	"fmt"

	_ "github.com/lib/pq"
	"net/http"
	"ppio/routes"
)

/**
Function which inserts dummy data into the database.
*/
func initialiseDb() *sql.DB {

	db, err := sql.Open("postgres", "postgresql://ppio_user@localhost:26257/ppio?sslmode=disable")
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	return db
}

func fillDb(dbConn *sql.DB) {

	players := utils.GetPlayers()

	for _, player := range players {
		lastID := player.Insert(dbConn)
		player.ID = lastID
	}
	fmt.Println("Players inserted")

	games := utils.GenerateGames(players)

	for _, game := range games {

		_ = game.Insert(dbConn)
	}
	fmt.Println("Games inserted")
}

func main() {

	end := make(chan bool)

	dbConn := initialiseDb()
	defer dbConn.Close()

	if len(os.Args) == 2 && os.Args[1] == "fillDb" {
		fillDb(dbConn)
	}


	// Handle the routes with gorillamux

	go func() {
		http.ListenAndServe(":9000", routes.GetRouter(dbConn))
	}()


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
