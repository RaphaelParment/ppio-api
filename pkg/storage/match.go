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
		return nil, fmt.Errorf("failed to query for all matches")
	}

	for rows.Next() {
		var m core.Match
		err := rows.Scan(&m.ID, &m.HomePlayerID, &m.AwayPlayerID, &m.HomePlayerScore, &m.AwayPlayerScore, &m.Datetime)
		if err != nil {
			return nil, fmt.Errorf("failed to scan match")
		}
		matches = append(matches, &m)
	}
	return matches, nil
}
