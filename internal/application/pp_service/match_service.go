package pp_service

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	matchRepository "github.com/RaphaelParment/ppio-api/internal/domain/match/repository"
)

type matchService struct {
	findOnePersister matchRepository.FinderOnePersister
}

func NewMatchService(findOnePersister matchRepository.FinderOnePersister) *matchService {
	return &matchService{findOnePersister: findOnePersister}
}

func (s *matchService) HandleFindMatch(ctx context.Context, id matchModel.Id) (matchModel.Match, error) {
	return s.findOnePersister.FindOne(ctx, id)
}
