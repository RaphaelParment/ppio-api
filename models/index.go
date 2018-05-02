package models

// DefaultLimit Default limit used in queries
const DefaultLimit = 20

// DefaultOffset Default offset used in queries
const DefaultOffset = 0

const startQueryCount = "SELECT COUNT(0) FROM "
const startGameQuery = "SELECT g.id, g.player1_id, g.player2_id, g.winner_id, g.validation_state, g.edited_by_id, g.datetime FROM "
const startPlayerQuery = "SELECT p.id, p.first_name, p.last_name, p.email, p.points FROM "
