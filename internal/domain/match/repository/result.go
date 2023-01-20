package repository

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
)

type FindResult interface {
	Find(ctx context.Context, id matchModel.Id) (matchModel.Result, error)
}
