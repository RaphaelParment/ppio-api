package repository

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type FindMatch interface {
	Find(ctx context.Context, id matchModel.Id) (matchModel.Match, error)
}

type FindAllMatches interface {
	FindAll(ctx context.Context) ([]matchModel.Match, error)
}

type PersistOneMatch interface {
	Persist(ctx context.Context, playerOneId, playerTwoId playerModel.Id, matchTime time.Time) (matchModel.Match, error)
}

type FinderPersister interface {
	FindMatch
	FindAllMatches
	PersistOneMatch
}
