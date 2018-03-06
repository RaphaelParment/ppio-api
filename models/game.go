package models

import (
	"database/sql"
	"log"
	"time"
)

type Game struct {
	ID       int64     `json:"id"`
	DateTime time.Time `json:"datetime"`
	Player1  Player    `json:"player1"`
	Player2  Player    `json:"player2"`
	Score1   int       `json:"score1"`
	Score2   int       `json:"score2"`
}

func (game *Game) Insert(dbConn *sql.DB) int64 {

	var id int64
	err := dbConn.QueryRow("INSERT INTO game (player1_id, player2_id, score1, score2, datetime) VALUES($1, $2, $3, $4, $5) RETURNING id", game.Player1.ID, game.Player2.ID, game.Score1, game.Score2, game.DateTime).Scan(&id)

	if err != nil {
		log.Fatalf("Could not insert game %v. Error: %v\n", game, err)
	}

	return id
}
