package rest

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"log"
	"time"
)

type MatchService interface {
	HandleFindMatch(ctx context.Context, id matchModel.Id) (matchModel.Match, error)
	HandleFindMatches(ctx context.Context) ([]matchModel.Match, error)
	HandlePersistMatch(ctx context.Context, playerOneId, playerTwoId playerModel.Id, datetime time.Time) (matchModel.Match, error)
}

type server struct {
	logger       *log.Logger
	matchService MatchService
}

func NewServer(logger *log.Logger, matchService MatchService) *server {
	return &server{logger: logger, matchService: matchService}
}
