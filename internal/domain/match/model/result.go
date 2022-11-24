package model

import playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"

type Result struct {
	winnerID     playerModel.Id
	loserRetired bool
}

func NewResult(winnerID playerModel.Id, loserRetired bool) Result {
	return Result{winnerID: winnerID, loserRetired: loserRetired}
}

func (r Result) WinnerID() playerModel.Id {
	return r.winnerID
}

func (r Result) LoserRetired() bool {
	return r.loserRetired
}
