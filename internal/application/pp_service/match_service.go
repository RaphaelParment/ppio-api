package pp_service

import (
	"context"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	matchRepository "github.com/RaphaelParment/ppio-api/internal/domain/match/repository"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type matchValidator interface {
	ValidateMatch(match matchModel.Match) map[string]string
}

type matchService struct {
	finderPersister matchRepository.FinderPersisterUpdater
	matchValidator  matchValidator
}

func NewMatchService(findOnePersister matchRepository.FinderPersisterUpdater, matchValidator matchValidator) *matchService {
	return &matchService{finderPersister: findOnePersister, matchValidator: matchValidator}
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

func (s *matchService) HandleUpdateOneMatch(ctx context.Context, id matchModel.Id, matchPatch matchModel.MatchPatch) (matchModel.Match, error) {
	var (
		playerOneId playerModel.Id
		playerTwoId playerModel.Id
		result      matchModel.Result
		score       matchModel.Score
		datetime    time.Time
	)

	// load the current match
	match, err := s.finderPersister.Find(ctx, id)
	if err != nil {
		return matchModel.Match{}, err
	}

	if matchPatch.PlayerOneId() != nil {
		playerOneId = *matchPatch.PlayerOneId()
	} else {
		playerOneId = match.PlayerOneId()
	}

	if matchPatch.PlayerTwoId() != nil {
		playerTwoId = *matchPatch.PlayerTwoId()
	} else {
		playerTwoId = match.PlayerTwoId()
	}

	if matchPatch.Result() != nil {
		result = *matchPatch.Result()
	} else {
		result = match.Result()
	}

	if matchPatch.Score() != nil {
		score = *matchPatch.Score()
	} else {
		score = match.Score()
	}

	if matchPatch.Datetime() != nil {
		datetime = *matchPatch.Datetime()
	} else {
		datetime = match.Datetime()
	}

	// update current match in memory
	patchedMatch := matchModel.NewMatch(
		match.Id(),
		playerOneId,
		playerTwoId,
		result,
		score,
		datetime,
	)

	// validate that new match is still valid
	problems := s.matchValidator.ValidateMatch(patchedMatch)
	if len(problems) > 0 {
		// TODO problems to err
		return matchModel.Match{}, nil
	}

	_, err = s.finderPersister.Update(ctx, patchedMatch)
	if err != nil {
		return matchModel.Match{}, err
	}

	return patchedMatch, nil
}
