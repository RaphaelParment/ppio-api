package entity

import (
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type MatchPatch struct {
	PlayerOneId *int    `json:"player_one_id"`
	PlayerTwoId *int    `json:"player_two_id"`
	Result      *Result `json:"result"`
	Score       *Score  `json:"score"`
	Datetime    *string `json:"datetime"`
}

func MatchPatchFromJSON(matchPatch MatchPatch) (matchModel.MatchPatch, error) {
	var (
		playerOneId *playerModel.Id
		playerTwoId *playerModel.Id
		result      *matchModel.Result
		score       *matchModel.Score
		datetime    *time.Time
	)

	if matchPatch.PlayerOneId != nil {
		tempPlayerId := playerModel.Id(*matchPatch.PlayerOneId)
		playerOneId = &tempPlayerId
	}

	if matchPatch.PlayerTwoId != nil {
		tempPlayerId := playerModel.Id(*matchPatch.PlayerTwoId)
		playerTwoId = &tempPlayerId
	}

	if matchPatch.Result != nil {
		tempResult := matchModel.NewResult(
			playerModel.Id(matchPatch.Result.WinnerID),
			matchPatch.Result.LoserRetired,
		)
		result = &tempResult
	}

	if matchPatch.Score != nil {
		var tempScore matchModel.Score
		for _, set := range *matchPatch.Score {
			tempScore = append(tempScore, matchModel.NewSet(set.PlayerOneScore, set.PlayerTwoScore))
		}
		score = &tempScore
	}

	if matchPatch.Datetime != nil {
		tempDatetime, err := time.Parse(time.DateTime, *matchPatch.Datetime)
		if err != nil {
			return matchModel.MatchPatch{}, err
		}
		datetime = &tempDatetime
	}

	return matchModel.NewMatchPatch(
		playerOneId,
		playerTwoId,
		result,
		score,
		datetime,
	), nil
}
