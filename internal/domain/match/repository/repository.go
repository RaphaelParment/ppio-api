package repository

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type FindOne interface {
	FindOne(ctx context.Context, id matchModel.Id) (matchModel.Game, error)
}

type FindAll interface {
	FindAll(ctx context.Context) ([]matchModel.Game, error)
}

type Persister interface {
	Persist(ctx context.Context, playerOneId, playerTwoId playerModel.Id, matchTime time.Time) (matchModel.Game, error)
}

type FinderPersister interface {
	FindOne
	FindAll
	Persister
}
