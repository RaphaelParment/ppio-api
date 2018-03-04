package database

import (
	"gopkg.in/olivere/elastic.v5"
	"context"
)

type IDbObject interface {
	Insert(client *elastic.Client, ctx context.Context, id int)
}
