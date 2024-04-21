package entity

import (
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"time"
)

type Match struct {
	Id          int    `json:"id"`
	PlayerOneId int    `json:"player_one_id"`
	PlayerTwoId int    `json:"player_two_id"`
	Result      Result `json:"result"`
	Score       Score  `json:"score"`
	Datetime    string `json:"datetime"`
}

func MatchToJSON(match matchModel.Match) Match {
	var score Score
	for _, set := range match.Score() {
		score = append(score, Set{
			PlayerOneScore: set.PlayerOneScore(),
			PlayerTwoScore: set.PlayerTwoScore(),
		})
	}

	return Match{
		Id:          match.Id().Int(),
		PlayerOneId: match.PlayerOneId().AsInt(),
		PlayerTwoId: match.PlayerTwoId().AsInt(),
		Result: Result{
			WinnerID:     match.Result().WinnerID().AsInt(),
			LoserRetired: match.Result().LoserRetired(),
		},
		Score:    score,
		Datetime: match.Datetime().Format(time.DateTime),
	}
}

func MatchFromJSON(match Match) (matchModel.Match, error) {
	var score matchModel.Score

	for _, set := range match.Score {
		score = append(score, matchModel.NewSet(set.PlayerOneScore, set.PlayerTwoScore))
	}

	datetime, err := time.Parse(time.DateTime, match.Datetime)
	if err != nil {
		return matchModel.Match{}, err
	}

	return matchModel.NewMatch(
		matchModel.Id(match.Id),
		playerModel.Id(match.PlayerOneId),
		playerModel.Id(match.PlayerTwoId),
		matchModel.NewResult(
			playerModel.Id(match.Result.WinnerID),
			match.Result.LoserRetired,
		),
		score,
		datetime,
	), nil
}
