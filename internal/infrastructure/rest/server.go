package rest

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"log"
	"time"
)

type GameService interface {
	HandleFindGame(ctx context.Context, id matchModel.Id) (matchModel.Game, error)
	HandleFindGames(ctx context.Context) ([]matchModel.Game, error)
	HandlePersistGame(ctx context.Context, playerOneId, playerTwoId playerModel.Id, datetime time.Time) (matchModel.Game, error)
}

type server struct {
	logger      *log.Logger
	gameService GameService
}

func NewServer(logger *log.Logger, gameService GameService) *server {
	return &server{logger: logger, gameService: gameService}
}
