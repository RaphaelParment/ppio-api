package models

import (
	"database/sql"
	"log"
)

// Set structure
type Set struct {
	ID     string `json:"id,omitempty"`
	GameID string `json:"gameId,omitempty"`
	Score1 int    `json:"score1"`
	Score2 int    `json:"score2"`
}

// Insert function
func (set *Set) Insert(dbConn *sql.DB) (int64, error) {

	var id int64
	err := dbConn.QueryRow(
		`INSERT INTO set (game_id, score1, score2) 
		VALUES($1, $2, $3) RETURNING id`,
		set.GameID, set.Score1, set.Score2).Scan(&id)

	if err != nil {
		log.Printf("Could not insert set %v. Error: %v\n", set, err)
		return 0, err
	}

	log.Printf("Inserted set with ID '%d'", id)

	return id, nil
}

// GetByID Returns a set based on its id
func (set *Set) GetByID(dbConn *sql.DB) error {

	err := dbConn.QueryRow(`
		SELECT id, score1, score2 WHERE id = $1`,
		&set.ID).
		Scan(&set.ID, &set.Score1, &set.Score2)

	if err != nil {
		log.Printf("Could not get set %v, err: %v", set, err)
		return err
	}

	return nil
}
