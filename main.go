package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RaphaelParment/ppio-api/app"
	"github.com/RaphaelParment/ppio-api/database"
	"github.com/RaphaelParment/ppio-api/handlers"

	_ "github.com/lib/pq"
)

func main() {
	logger := log.New(os.Stdout, "ppio: ", log.LstdFlags)

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	app := app.App{}
	app.Logger = logger
	app.CreateRouter()
	app.Database = db

	// Creating Players handler
	app.Players = handlers.NewPlayers()

	logger.Print("main : Listening :9001")
	log.Fatal(http.ListenAndServe(":9001", app.Router))
}
