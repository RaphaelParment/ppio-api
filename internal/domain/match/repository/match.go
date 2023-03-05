package repository

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
)

type FindMatch interface {
	Find(ctx context.Context, id matchModel.Id) (matchModel.Match, error)
}

type FindAllMatches interface {
	FindAll(ctx context.Context) ([]matchModel.Match, error)
}

type PersistOneMatch interface {
	Persist(ctx context.Context, match matchModel.Match) (matchModel.Id, error)
}

type FinderPersister interface {
	FindMatch
	FindAllMatches
	PersistOneMatch
}
