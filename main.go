package main

import (
	"database/sql"
	"log"
	"os"
	"ppio/utils"

	"fmt"

	"flag"
	"net/http"
	"ppio/routes"

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

	return db
}

func fillDb(dbConn *sql.DB) {

	players := utils.GetPlayers()

	for _, player := range players {
		lastID, _ := player.Insert(dbConn)
		player.ID = lastID
	}
	fmt.Println("Players inserted")

	games := utils.GenerateGames(players)

	for _, game := range games {

		_, _ = game.Insert(dbConn)
	}
	fmt.Println("Games inserted")
}

func main() {

	initDbData := flag.Bool("initDbData", false, "Insert dummy data in database.")

	flag.Parse()
	end := make(chan bool)

	dbConn := initialiseDb()
	defer dbConn.Close()

	if *initDbData {
		fillDb(dbConn)
		log.Printf("Initialised the database with dummy data. Terminating...")
		os.Exit(0)
	}

	// Handle the routes with gorillamux

	go func() {
		http.ListenAndServe(":9001", routes.GetRouter(dbConn))
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