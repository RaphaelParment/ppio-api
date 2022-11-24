package entity

import matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"

type Result struct {
	MatchID      int  `json:"match_id"`
	WinnerID     int  `json:"winner_id"`
	LoserRetired bool `json:"loser_retired"`
}

func MatchResultToJSON(result matchModel.Result) Result {
	return Result{
		WinnerID:     result.WinnerID().AsInt(),
		LoserRetired: result.LoserRetired(),
	}
}
