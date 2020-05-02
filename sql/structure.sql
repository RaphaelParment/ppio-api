CREATE DATABASE ppio;

CREATE TABLE player (
    id serial PRIMARY KEY,
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    email VARCHAR(64) NOT NULL UNIQUE,
    points SMALLINT
);

CREATE TABLE match (
    id serial PRIMARY KEY,
    first_player_id INTEGER REFERENCES player(id),
    second_player_id INTEGER REFERENCES player(id),
    date_time TIMESTAMP
);

CREATE TABLE match_result (
    match_id INTEGER REFERENCES match(id),
    winner_id INTEGER REFERENCES player(id),
    games_played SMALLINT,
    loser_retired BOOLEAN
);

CREATE TABLE game_score (
    match_id INTEGER REFERENCES match(id),
    game_number SMALLINT,
    first_player_score SMALLINT,
    second_player_score SMALLINT
);