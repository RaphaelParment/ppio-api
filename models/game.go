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

func (game *Game) Insert(dbConn *sql.DB) (int64, error) {

	var id int64
	err := dbConn.QueryRow("INSERT INTO game (player1_id, player2_id, score1, score2, datetime) VALUES($1, $2, $3, $4, $5) RETURNING id", game.Player1ID, game.Player2ID, game.Score1, game.Score2, game.DateTime).Scan(&id)

	if err != nil {
		log.Printf("Could not insert game %v. Error: %v\n", game, err)
		return 0, err
	}

	log.Printf("Inserted game with ID '%d'", id)

	return id, nil
}

func (game *Game) GetByID(dbConn *sql.DB) error {

	err := dbConn.QueryRow("SELECT id, player1_id, player2_id, score1, score2, datetime FROM game WHERE id = $1",
		&game.ID).Scan(&game.ID, &game.Player1ID, &game.Player2ID, &game.Score1, &game.Score2, &game.DateTime)

	if err != nil {
		log.Printf("Could not get game %v, err: %v", game, err)
		return err
	}

	return nil
}

func (game *Game) GetAll(dbConn *sql.DB) ([]Game, error) {

	games := make([]Game, 0, 512)
	rows, err := dbConn.Query("SELECT id, player1_id, player2_id, score1, score2, datetime FROM game")

	if err != nil {
		log.Printf("Could not fetch all games in DB. Error: %v", err)
		return nil, err
	}

	for rows.Next() {
		var game Game
		rows.Scan(&game.ID, &game.Player1ID, &game.Player2ID, &game.Score1, &game.Score2, &game.DateTime)
		games = append(games, game)
	}

	return games, nil
}

func (game *Game) Update(dbConn *sql.DB) (int64, error) {

	result, err := dbConn.Exec("UPDATE game SET player1_id = $1, player2_id = $2, score1 = $3, score2 = $4, datetime = $5 WHERE id = $6",
		game.Player1ID, game.Player2ID, game.Score1, game.Score2, game.DateTime, game.ID)

	if err != nil {
		log.Printf("Could not update game: %v, err: %v", game, err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Could not get the amount of affected rows. Error: %v", err)
		return 0, err
	}

	return rowsAffected, nil

}
