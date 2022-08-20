package persistence

//// GetMatchGamesScores returns the scores of all games of a given match
//func GetMatchGamesScores(db *sql.DB, id int) ([]*match_score.GameScore, error) {
//	var scores []*match_score.GameScore
//	q := "SELECT * FROM game_score where match_id = $1"
//	rows, err := db.Query(q, id)
//	if err != nil {
//		return nil, fmt.Errorf("could not query match %d games scores; %v", id, err)
//	}
//	defer rows.Close()
//	for rows.Next() {
//		var s match_score.GameScore
//		err := rows.Scan(&s.MatchID, &s.GameNbr, &s.FirstPlayerScore, &s.SecondPlayerScore)
//		if err != nil {
//			return nil, fmt.Errorf("could not scan game score; %v", err)
//		}
//		scores = append(scores, &s)
//	}
//	return scores, nil
//}
//
//// AddMatchGamesScores adds the games scores of a match
//func AddMatchGamesScores(db *sql.DB, scores match.GameScores) error {
//	nbrScores := len(scores)
//	q := "INSERT INTO game_score (match_id, game_number, first_player_score, second_player_score) VALUES "
//	for i := 0; i < nbrScores-1; i++ {
//		q += fmt.Sprintf("(%d, %d, %d, %d),", scores[i].MatchID, scores[i].GameNbr,
//			scores[i].FirstPlayerScore, scores[i].SecondPlayerScore)
//	}
//	q += fmt.Sprintf("(%d, %d, %d, %d)", scores[nbrScores-1].MatchID, scores[nbrScores-1].GameNbr,
//		scores[nbrScores-1].FirstPlayerScore, scores[nbrScores-1].SecondPlayerScore)
//
//	_, err := db.Exec(q)
//	if err != nil {
//		return fmt.Errorf("could not add match games scores; %v", err)
//	}
//
//	return nil
//}
