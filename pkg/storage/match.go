package storage

import (
	"database/sql"
	"fmt"

	"github.com/RaphaelParment/ppio-api/pkg/core"
)

// GetMatches returns all matches
func GetMatches(db *sql.DB) (core.Matches, error) {
	var matches core.Matches
	rows, err := db.Query("SELECT * FROM match")
	defer rows.Close()
	if err != nil {
		return nil, fmt.Errorf("could not query all matches; %v", err)
	}

	for rows.Next() {
		var m core.Match
		err := rows.Scan(&m.ID, &m.FirstPlayerID, &m.SecondPlayerID, &m.Datetime)
		if err != nil {
			return nil, fmt.Errorf("could not scan match; %v", err)
		}
		matches = append(matches, &m)
	}
	return matches, nil
}

// AddMatch adds a match
func AddMatch(db *sql.DB, m *core.Match) error {
	q := "INSERT INTO match (first_player_id, second_player_id, date_time) VALUES ($1, $2, $3)"
	_, err := db.Exec(q, m.FirstPlayerID, m.SecondPlayerID, m.Datetime)
	if err != nil {
		return fmt.Errorf("could not add match; %v", err)
	}

	return nil
}
