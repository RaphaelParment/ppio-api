package model

type Result struct {
	winnerID     int
	loserRetired bool
}

func NewResult(winnerID int, loserRetired bool) Result {
	return Result{winnerID: winnerID, loserRetired: loserRetired}
}

func (r Result) WinnerID() int {
	return r.winnerID
}

func (r Result) LoserRetired() bool {
	return r.loserRetired
}
