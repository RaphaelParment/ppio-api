package models

import (
	"gopkg.in/olivere/elastic.v5"
	"context"
	"strconv"
)

type Game struct {
	DateTime string `json:"datetime"`
	Player1  Player `json:"player1"`
	Player2  Player `json:"player2"`
	Score1   int    `json:"score1"`
	Score2   int    `json:"score2"`
}

func (game Game) Insert(client *elastic.Client, ctx context.Context, id int) {
	_, err := client.Index().
		Index("games").
		Type("doc").
		Id(strconv.Itoa(id)).
		BodyJson(game).
		Do(ctx)

	if err != nil {
		panic(err)
	}
}
