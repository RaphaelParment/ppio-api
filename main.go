package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"ppio-web/routes"
	"ppio-web/utils"
//	"time"

//	"github.com/coreos/go-systemd/daemon"
	"gopkg.in/olivere/elastic.v5"
	"fmt"

)


/**
Function which inserts dummy data into the database.
 */
func initialiseDb(ctx context.Context, client *elastic.Client) {

	players := utils.GetPlayers()

	for i, player := range players {
		player.Insert(client, ctx, i)
	}
	fmt.Println("Players inserted")

	games := utils.GenerateGames(players)

	for j, game := range games {

		game.Insert(client, ctx, j)
	}
	fmt.Println("Games inserted")
}


func main() {

	end := make(chan bool)
	//ctx := context.Background()

	client, err := elastic.NewClient(
                elastic.SetURL("http://172.17.0.2:9200"),                           // Docker default public address for elasticsearch.
                elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)), // Specific logger for the package.
        )
        if err != nil {
                log.Fatalf("Could not connect to the ElasticSearch instance: %v\n", err)
        }
        defer client.Stop()

	//initialiseDb(ctx, client)

	// Handle the routes with gorillamux
	go func() {
		http.ListenAndServe(":9000", routes.GetRouter(client))
	}()

	//go func() {
	//	interval, err := daemon.SdWatchdogEnabled(false)
	//	if err != nil || interval == 0 {
	//		log.Printf("Could not start the watchdog for SystemD...")
	//	}
	//	for {
	//		daemon.SdNotify(false, "WATCHDOG=1")
	//		time.Sleep(interval / 3)
	//	}
	//}()

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
