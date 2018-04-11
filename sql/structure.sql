CREATE TABLE game (
    id INT NOT NULL DEFAULT unique_rowid(),
    player1_id INTEGER NOT NULL,
    player2_id INTEGER NOT NULL,
    datetime TIMESTAMP WITH TIME ZONE NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    CONSTRAINT player1_fk FOREIGN KEY (player1_id) REFERENCES player (id),
    INDEX game_player1_id_idx (player1_id ASC),
    CONSTRAINT player2_fk FOREIGN KEY (player2_id) REFERENCES player (id),
    INDEX game_player2_id_idx (player2_id ASC),
    FAMILY "primary" (id, player1_id, player2_id, datetime)
);

CREATE TABLE set (
    id INT NOT NULL DEFAULT unique_rowid(),
    game_id INT NOT NULL,
    score1 INT NOT NULL DEFAULT 0,
    score2 INT NOT NULL DEFAULT 0,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    CONSTRAINT game_id_fk FOREIGN KEY (game_id) REFERENCES game(id),
    FAMILY "primary" (id, game_id, score1, score2)
);

CREATE TABLE player (
    id INT NOT NULL DEFAULT unique_rowid(),
    first_name STRING NULL,
    last_name STRING NULL,
    points INT NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    FAMILY "primary" (id, first_name, last_name, points)
);