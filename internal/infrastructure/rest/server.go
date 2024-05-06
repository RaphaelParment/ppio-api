package rest

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	"log"
)

type MatchService interface {
	HandleFindMatch(ctx context.Context, id matchModel.Id) (matchModel.Match, error)
	HandleFindMatches(ctx context.Context) ([]matchModel.Match, error)
	HandlePersistMatch(ctx context.Context, match matchModel.Match) (matchModel.Id, error)
	HandleUpdateOneMatch(ctx context.Context, id matchModel.Id, matchPatch matchModel.MatchPatch) (matchModel.Match, error)
}

type server struct {
	logger       *log.Logger
	matchService MatchService
}

func NewServer(logger *log.Logger, matchService MatchService) *server {
	return &server{logger: logger, matchService: matchService}
}
