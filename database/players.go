package database

import (
	"database/sql"
	"fmt"

	"github.com/RaphaelParment/ppio-api/data"
)

// GetPlayers returns all players
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

// GetPlayer returns the player with id <id>
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

// AddPlayer inserts the player pointed to by <p>
func AddPlayer(db *sql.DB, p *data.Player) error {
	q := "INSERT INTO player (first_name, last_name, email, points) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(q, p.FirstName, p.LastName, p.Email, p.Points)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePlayer updates the player with id <id> with the field of the player pointed to by <p>.
// If there is no player with id <id>, false is returned as first parameter
func UpdatePlayer(db *sql.DB, id int, p *data.Player) (bool, error) {
	q := "UPDATE player SET first_name = $1, last_name = $2, email = $3, points = $4 WHERE id = $5"
	res, err := db.Exec(q, p.FirstName, p.LastName, p.Email, p.Points, id)
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

// DeletePlayer removes the player with id <id>. If no player with id <id> is
// found in the database, false is returned as first paramater
func DeletePlayer(db *sql.DB, id int) (bool, error) {
	q := "DELETE FROM player WHERE id = $1"
	res, err := db.Exec(q, id)
	if err != nil {
		return false, err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
