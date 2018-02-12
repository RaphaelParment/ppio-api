package main

import (
	"gopkg.in/olivere/elastic.v5"
	"context"
	"ppio-web/src/models"
	"ppio-web/src/config"
	"ppio-web/src/utils"
	"strconv"
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

	_, err = client.CreateIndex(indexName).
		BodyString(config.PlayerMapping).Do(ctx)
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

	createIndex("players", client, ctx)
	createIndex("games", client, ctx)

	var players []models.Player
	players = utils.GetPlayers()

	for index, player := range players {
		_, err = client.Index().
			Index("players").
			Type("doc").
			Id(strconv.Itoa(index)).
			BodyJson(player).
			Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}