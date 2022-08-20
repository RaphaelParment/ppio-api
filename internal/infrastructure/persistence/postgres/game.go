package postgres

import (
	"context"
	"database/sql"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence/postgres/entity"
	"log"
	"time"
)

type matchStore struct {
	logger *log.Logger
	db     *sql.DB
}

func NewMatchStore(logger *log.Logger, db *sql.DB) *matchStore {
	return &matchStore{logger: logger, db: db}
}

func (s *matchStore) FindOne(ctx context.Context, id matchModel.Id) (matchModel.Game, error) {
	var match matchModel.Game
	row := s.db.QueryRowContext(ctx, "SELECT * FROM match WHERE id = $1", id)
	err := row.Scan(&match.Id, &match.PlayerOneId, &match.PlayerTwoId, &match.Datetime)
	if err != nil {
		return matchModel.Game{}, err
	}

	return match, nil
}

func (s *matchStore) FindAll(ctx context.Context) ([]matchModel.Game, error) {
	var matches []matchModel.Game
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM match")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()

	var match entity.Match
	for rows.Next() {
		err := rows.Scan(&match.Id, &match.PlayerOneId, &match.PlayerTwoId, &match.Datetime)
		if err != nil {
			return nil, err
		}

		matches = append(matches, entity.MatchFromJSON(match))
	}

	return matches, nil
}

func (s *matchStore) Persist(
	ctx context.Context,
	playerOneId playerModel.Id,
	playerTwoId playerModel.Id,
	matchTime time.Time,
) (matchModel.Game, error) {
	query := "INSERT INTO match (first_player_id, second_player_id, date_time) VALUES ($1, $2, $3) RETURNING id"

	var id int
	err := s.db.QueryRowContext(ctx, query, int32(playerOneId), int32(playerTwoId), matchTime).Scan(&id)
	if err != nil {
		s.logger.Printf("failed to insert match; %s", err)
		return matchModel.Game{}, err
	}

	match := matchModel.Game{
		Id:          matchModel.Id(id),
		PlayerOneId: playerOneId,
		PlayerTwoId: playerTwoId,
		Datetime:    matchTime,
	}

	return match, nil
}
