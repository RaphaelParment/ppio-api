package pp_service

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	matchRepository "github.com/RaphaelParment/ppio-api/internal/domain/match/repository"
)

type matchService struct {
	finderPersister matchRepository.FinderPersister
}

func NewMatchService(findOnePersister matchRepository.FinderPersister) *matchService {
	return &matchService{finderPersister: findOnePersister}
}

func (s *matchService) HandleFindMatch(ctx context.Context, id matchModel.Id) (matchModel.Match, error) {
	return s.finderPersister.Find(ctx, id)
}

func (s *matchService) HandleFindMatches(ctx context.Context) ([]matchModel.Match, error) {
	return s.finderPersister.FindAll(ctx)
}

func (s *matchService) HandlePersistMatch(ctx context.Context, match matchModel.Match) (matchModel.Id, error) {
	return s.finderPersister.Persist(ctx, match)
}
