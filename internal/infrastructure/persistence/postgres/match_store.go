package postgres

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence/postgres/entity"
	"github.com/jmoiron/sqlx"
	"log"
)

type matchStore struct {
	logger *log.Logger
	db     *sqlx.DB
}

func NewMatchStore(logger *log.Logger, db *sqlx.DB) *matchStore {
	return &matchStore{logger: logger, db: db}
}

func (s *matchStore) Find(ctx context.Context, id matchModel.Id) (matchModel.Match, error) {
	return s.find(ctx, id)
}

func (s *matchStore) FindAll(ctx context.Context) ([]matchModel.Match, error) {
	var matchIds []matchModel.Id
	rowsId, err := s.db.QueryxContext(ctx, "SELECT id FROM match")
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := rowsId.Close(); err != nil {
			s.logger.Printf("failed to close rows id; %v", err)
		}
	}()

	var id int
	for rowsId.Next() {
		err = rowsId.Scan(&id)
		if err != nil {
			return nil, err
		}

		matchIds = append(matchIds, matchModel.Id(id))
	}

	var matches []matchModel.Match
	for _, matchId := range matchIds {
		match, err := s.find(ctx, matchId)
		if err != nil {
			return nil, err
		}

		matches = append(matches, match)
	}

	return matches, nil
}

func (s *matchStore) Persist(ctx context.Context, match matchModel.Match) (matchModel.Id, error) {
	var matchId int64

	tx, err := s.db.Beginx()
	if err != nil {
		return matchModel.NewUndefinedId(), err
	}

	err = tx.QueryRowxContext(
		ctx,
		"INSERT INTO match (player_one_id, player_two_id, date_time) VALUES ($1, $2, $3) RETURNING id",
		match.PlayerOneId().AsInt(),
		match.PlayerTwoId().AsInt(),
		match.Datetime(),
	).Scan(&matchId)
	if err != nil {
		return matchModel.NewUndefinedId(), err
	}

	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO match_result(match_id, winner_id, loser_retired) VALUES ($1, $2, $3)",
		matchId,
		match.Result().WinnerID(),
		match.Result().LoserRetired(),
	)
	if err != nil {
		return matchModel.NewUndefinedId(), err
	}

	for _, set := range match.Score() {
		_, err = tx.ExecContext(
			ctx,
			"INSERT INTO set (player_one_score, player_two_score) VALUES ($1, $2)",
			set.PlayerOneScore(),
			set.PlayerTwoScore(),
		)
		if err != nil {
			return matchModel.NewUndefinedId(), err
		}
	}

	err = tx.Commit()
	if err != nil {
		return matchModel.NewUndefinedId(), err
	}

	return matchModel.Id(matchId), nil
}

func (s *matchStore) Update(ctx context.Context, match matchModel.Match) (matchModel.Id, error) {
	_, err := s.db.ExecContext(
		ctx,
		"UPDATE match SET player_one_id = $1, player_two_id = $2 WHERE id = $3",
		match.PlayerOneId().AsInt(),
		match.PlayerTwoId().AsInt(),
		match.Id().Int(),
	)
	if err != nil {
		return matchModel.NewUndefinedId(), err
	}

	return match.Id(), nil
}

func (s *matchStore) find(ctx context.Context, id matchModel.Id) (matchModel.Match, error) {
	var (
		match  entity.Match
		result entity.MatchResult
	)

	rowMatch := s.db.QueryRowxContext(
		ctx,
		"SELECT id, player_one_id, player_two_id, date_time FROM match WHERE id = $1",
		id,
	)
	err := rowMatch.StructScan(&match)
	if err != nil {
		return matchModel.Match{}, err
	}

	rowResult := s.db.QueryRowxContext(
		ctx,
		"SELECT match_id, winner_id, loser_retired FROM match_result WHERE match_id = $1",
		id,
	)
	err = rowResult.StructScan(&result)
	if err != nil {
		return matchModel.Match{}, err
	}

	rowsScore, err := s.db.QueryxContext(
		ctx,
		`SELECT id, player_one_score, player_two_score 
FROM set s 
JOIN match_sets ms 
	ON s.id = ms.set_id
WHERE ms.match_id = $1`,
		id,
	)
	if err != nil {
		return matchModel.Match{}, err
	}
	defer func() {
		if err := rowsScore.Close(); err != nil {
			s.logger.Printf("failed to close rows score; %v", err)
		}
	}()

	var (
		set   entity.Set
		score matchModel.Score
	)
	for rowsScore.Next() {
		err = rowsScore.StructScan(&set)
		if err != nil {
			return matchModel.Match{}, err
		}

		score = append(score, matchModel.NewSet(set.PlayerOneScore, set.PlayerTwoScore))
	}

	return matchModel.NewMatch(
		id,
		playerModel.Id(match.PlayerOneId),
		playerModel.Id(match.PlayerTwoId),
		matchModel.NewResult(
			playerModel.Id(result.WinnerID),
			result.LoserRetired,
		),
		score,
		match.Datetime,
	), nil
}
