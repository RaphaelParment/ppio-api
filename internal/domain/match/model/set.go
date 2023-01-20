package model

type Set struct {
	playerOneScore int
	playerTwoScore int
}

func NewSet(playerOneScore int, playerTwoScore int) Set {
	return Set{playerOneScore: playerOneScore, playerTwoScore: playerTwoScore}
}

func (s Set) PlayerOneScore() int {
	return s.playerOneScore
}

func (s Set) PlayerTwoScore() int {
	return s.playerTwoScore
}
