package postgres

import (
	"context"
	"database/sql"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type matchStore struct {
	db *sql.DB
}

func NewMatchStore(db *sql.DB) *matchStore {
	return &matchStore{db: db}
}

func (s *matchStore) FindOne(ctx context.Context, id matchModel.Id) (matchModel.Match, error) {
	var match matchModel.Match
	row := s.db.QueryRowContext(ctx, "SELECT * FROM match WHERE id = $1", id)
	err := row.Scan(&match.Id, &match.FirstPlayerId, &match.SecondPlayerId, &match.Datetime)
	if err != nil {
		return matchModel.Match{}, err
	}

	return match, nil
}

func (s *matchStore) Persist(
	ctx context.Context,
	firstPlayerId playerModel.Id,
	secondPlayerId playerModel.Id,
	matchTime time.Time,
) (matchModel.Match, error) {
	query := "INSERT INTO match (first_player_id, second_player_id, date_time) VALUES ($1, $2, $3)"
	result, err := s.db.ExecContext(ctx, query, firstPlayerId, secondPlayerId, matchTime)
	if err != nil {
		return matchModel.Match{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return matchModel.Match{}, err
	}

	match := matchModel.Match{
		Id:             matchModel.Id(id),
		FirstPlayerId:  firstPlayerId,
		SecondPlayerId: secondPlayerId,
		Datetime:       matchTime,
	}

	return match, nil
}
