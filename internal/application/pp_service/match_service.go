package pp_service

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	matchRepository "github.com/RaphaelParment/ppio-api/internal/domain/match/repository"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type gameService struct {
	finderPersister matchRepository.FinderPersister
}

func NewGameService(findOnePersister matchRepository.FinderPersister) *gameService {
	return &gameService{finderPersister: findOnePersister}
}

func (s *gameService) HandleFindGame(ctx context.Context, id matchModel.Id) (matchModel.Game, error) {
	return s.finderPersister.FindOne(ctx, id)
}

func (s *gameService) HandleFindGames(ctx context.Context) ([]matchModel.Game, error) {
	return s.finderPersister.FindAll(ctx)
}

func (s *gameService) HandlePersistGame(ctx context.Context, playerOneId, playerTwoId playerModel.Id, datetime time.Time) (matchModel.Game, error) {
	return s.finderPersister.Persist(ctx, playerOneId, playerTwoId, datetime)
}
