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

INSERT INTO player (id, first_name, last_name, email, points)
VALUES
    (1, 'Lucia', 'Moore', 'lucia.moore@ppio.com', 0),
    (2, 'Paula', 'Young', 'paula.young@ppio.com', 0),
    (3, 'Steven', 'Farmer', 'steven.farmer@ppio.com', 0),
    (4, 'Moses', 'Wong', 'moses.wong@ppio.com', 0),
    (5, 'Allen', 'Daniels', 'allen.daniels@ppio.com', 0),
    (6, 'claire', 'Warren', 'claire.warren@ppio.com', 0),
    (7, 'Conrad', 'Gibson', 'conrad.gibson@ppio.com', 0),
    (8, 'Ted', 'Little', 'ted.little@ppio.com', 0),
    (9, 'Jessie', 'Gibbs', 'jessie.gibbs@ppio.com', 0),
    (10, 'Ivan', 'Lee', 'ivan.lee@ppio.com', 0);