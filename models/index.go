package models

const DefaultLimit = 20
const DefaultOffset = 0

const startQueryCount = "SELECT COUNT(0) FROM "
const startGameQuery = "SELECT id, player1_id, player2_id, winner_id, validation_state, edited_by_id, datetime"
