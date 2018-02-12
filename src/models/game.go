package models

type Game struct {
	Date    string  `json:"date"`
	Player1 *Player `json:"player1"`
	Player2 *Player `json:"player2"`
	Score1  int     `json:"score1"`
	Score2  int     `json:"score2"`
}
