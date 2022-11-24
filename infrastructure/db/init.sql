CREATE TABLE player (
    id serial PRIMARY KEY,
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    email VARCHAR(64) NOT NULL UNIQUE,
    points SMALLINT
);

CREATE TABLE match (
    id serial PRIMARY KEY,
    player_one_id INTEGER REFERENCES player(id),
    player_two_id INTEGER REFERENCES player(id),
    date_time TIMESTAMP
);

CREATE TABLE set (
    id serial PRIMARY KEY,
    player_one_score SMALLINT,
    player_two_score SMALLINT
);

CREATE TABLE match_sets (
    match_id INTEGER REFERENCES match(id),
    set_id INTEGER REFERENCES set(id),
    CONSTRAINT match_set UNIQUE (match_id, set_id)
);

CREATE TABLE match_result (
    match_id INTEGER REFERENCES match(id),
    winner_id INTEGER REFERENCES player(id),
    loser_retired BOOLEAN
);

INSERT INTO player (id, first_name, last_name, email, points)
VALUES
    (1, 'Lucia', 'Moore', 'lucia.moore@ppio.com', 0),
    (2, 'Paula', 'Young', 'paula.young@ppio.com', 0),
    (3, 'Steven', 'Farmer', 'steven.farmer@ppio.com', 0),
    (4, 'Moses', 'Wong', 'moses.wong@ppio.com', 0),
    (5, 'Allen', 'Daniels', 'allen.daniels@ppio.com', 0),
    (6, 'Claire', 'Warren', 'claire.warren@ppio.com', 0),
    (7, 'Conrad', 'Gibson', 'conrad.gibson@ppio.com', 0),
    (8, 'Ted', 'Little', 'ted.little@ppio.com', 0),
    (9, 'Jessie', 'Gibbs', 'jessie.gibbs@ppio.com', 0),
    (10, 'Ivan', 'Lee', 'ivan.lee@ppio.com', 0);

SELECT setval('player_id_seq', 10, true);

INSERT INTO match(id, player_one_id, player_two_id, date_time)
VALUES
    (1, 1, 2, '2022-06-22 16:10:25-00'),
    (2, 3, 4, '2022-06-22 19:12:25-00'),
    (3, 5, 6, '2022-06-22 14:11:24-00'),
    (4, 7, 8, '2022-06-22 10:54:36-00'),
    (5, 9, 10, '2022-06-22 19:04:27-00');

SELECT setval('match_id_seq', 5, true);

INSERT INTO set (id, player_one_score, player_two_score)
VALUES
    (1, 11, 3),
    (2, 11, 9),
    (3, 8, 11),
    (4, 5, 11),
    (5, 4, 11),
    (6, 11, 9),
    (7, 11, 6),
    (8, 11, 2),
    (9, 11, 0),
    (10, 7, 11),
    (11, 8, 11);

SELECT setval('set_id_seq', 11, true);

INSERT INTO match_sets (match_id, set_id)
VALUES
    (1, 1),
    (1, 2),
    (2, 3),
    (2, 4),
    (3, 5),
    (3, 6),
    (3, 7),
    (4, 8),
    (4, 9),
    (5, 10),
    (5, 11);

INSERT INTO match_result (match_id, winner_id, loser_retired)
VALUES
    (1, 1, FALSE),
    (2, 4, FALSE),
    (3, 5, FALSE),
    (4, 7, FALSE),
    (5, 10, FALSE);