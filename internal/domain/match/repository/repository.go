package repository

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type FindOne interface {
	FindOne(ctx context.Context, id matchModel.Id) (matchModel.Match, error)
}

type Persister interface {
	Persist(
		ctx context.Context,
		firstPlayerId playerModel.Id,
		secondPlayerId playerModel.Id,
		matchTime time.Time,
	) (matchModel.Match, error)
}

type FinderOnePersister interface {
	FindOne
	Persister
}
