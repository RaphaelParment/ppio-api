package persistence

//// GetMatchResult returns the result of a given match
//func GetMatchResult(db *sql.DB, id int) (*match.MatchResult, error) {
//	var r match.MatchResult
//	q := "SELECT * FROM match_result WHERE match_id = $1"
//	row := db.QueryRow(q, id)
//	err := row.Scan(&r.MatchID, &r.WinnerID, &r.GamesPlayed, &r.LoserRetired)
//	if err == sql.ErrNoRows {
//		return nil, err
//	}
//	if err != nil {
//		return nil, fmt.Errorf("coult not to query / cast for match id: %d; %v", id, err)
//	}
//	return &r, nil
//}
//
//// AddMatchResult adds the result of a match
//func AddMatchResult(db *sql.DB, result *match.MatchResult) error {
//	q := "INSERT INTO match_result (match_id, winner_id, games_played, loser_retired) VALUES ($1, $2, $3, $4)"
//	_, err := db.Exec(q, result.MatchID, result.WinnerID, result.GamesPlayed, result.LoserRetired)
//	if err != nil {
//		return fmt.Errorf("could not add match result; %v", err)
//	}
//
//	return nil
//}
