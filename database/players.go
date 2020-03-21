package database

import (
	"database/sql"
	"fmt"

	"github.com/RaphaelParment/ppio-api/data"
)

// GetPlayers DB func
func GetPlayers(db *sql.DB) (data.Players, error) {
	var players data.Players
	rows, err := db.Query("SELECT * FROM player")
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to query for all players")
	}

	for rows.Next() {
		var p data.Player
		err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Email, &p.Points)
		if err != nil {
			return nil, fmt.Errorf("failed to scan players")
		}
		players = append(players, &p)
	}
	return players, nil
}

// GetPlayer returns the player with id specified in URL
func GetPlayer(db *sql.DB, id int) (*data.Player, error) {
	var p data.Player
	q := "SELECT * FROM player WHERE id = $1"
	row := db.QueryRow(q, id)
	err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Email, &p.Points)
	if err == sql.ErrNoRows {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query / cast for player id: %d", id)
	}
	return &p, nil
}
