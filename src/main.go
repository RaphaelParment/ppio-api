package main

import (
	"gopkg.in/olivere/elastic.v5"
	"context"
	"ppio/src/config"
	"ppio/src/utils"
)


func createIndex(indexName string, client *elastic.Client,
	ctx context.Context) {

	// Create the players index
	exists, err := client.IndexExists(indexName).Do(ctx)
	if exists {
		_, err := client.DeleteIndex(indexName).Do(ctx)
		if err != nil {
			panic(err)
		}
	}

	if indexName == "players" {
		_, err = client.CreateIndex(indexName).
			BodyString(config.PlayerMapping).Do(ctx)
	} else if indexName == "games" {
		_, err = client.CreateIndex(indexName).
			BodyString(config.GameMapping).Do(ctx)
	}

	if err != nil {
		panic(err)
	}
}




func main() {

	ctx := context.TODO()

	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Stop()


	players := utils.GetPlayers()

	for i, player := range players {
		player.Insert(client, ctx, i)
	}

	games := utils.GenerateGames(players)

	for j, game := range games {

		game.Insert(client, ctx, j)
	}

}