package models

import (
	"database/sql"
	"log"
	"time"
)

type Game struct {
	ID        int64     `json:"id,omitempty"`
	DateTime  time.Time `json:"datetime,omitempty"`
	Player1ID int64     `json:"player1Id"`
	Player2ID int64     `json:"player2Id"`
	Score1    int       `json:"score1"`
	Score2    int       `json:"score2"`
}

func (game *Game) Insert(dbConn *sql.DB) int64 {

	var id int64
	err := dbConn.QueryRow("INSERT INTO game (player1_id, player2_id, score1, score2, datetime) VALUES($1, $2, $3, $4, $5) RETURNING id", game.Player1ID, game.Player2ID, game.Score1, game.Score2, game.DateTime).Scan(&id)

	if err != nil {
		log.Fatalf("Could not insert game %v. Error: %v\n", game, err)
	}

	log.Printf("Inserted game with ID '%d'", id)

	return id
}
