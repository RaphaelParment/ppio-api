package repository

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
)

type FindSets interface {
	Find(ctx context.Context, id matchModel.Id) ([]matchModel.Set, error)
}
