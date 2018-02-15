package models

import (
	"gopkg.in/olivere/elastic.v5"
	"context"
	"strconv"
)


type Player struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Points    int    `json:"points"`
}

func (player Player) Insert(client *elastic.Client, ctx context.Context, id int) {
	_, err := client.Index().
		Index("players").
		Type("doc").
		Id(strconv.Itoa(id)).
		BodyJson(player).
		Do(ctx)

	if err != nil {
		panic(err)
	}
}