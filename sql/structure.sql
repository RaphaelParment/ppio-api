CREATE TABLE game (
id INT NOT NULL DEFAULT unique_rowid(),
player1_id INTEGER NOT NULL,
player2_id INTEGER NOT NULL,
score1 INT NULL,
score2 INT NULL,
datetime TIMESTAMP WITH TIME ZONE NULL,
CONSTRAINT "primary" PRIMARY KEY (id ASC),
CONSTRAINT player1_fk FOREIGN KEY (player1_id) REFERENCES player (id),
INDEX game_player1_id_idx (player1_id ASC),
CONSTRAINT player2_fk FOREIGN KEY (player2_id) REFERENCES player (id),
INDEX game_player2_id_idx (player2_id ASC),
FAMILY "primary" (id, player1_id, player2_id, score1, score2, datetime)
);

CREATE TABLE player (
id INT NOT NULL DEFAULT unique_rowid(),
first_name STRING NULL,
last_name STRING NULL,
score INT NULL,
CONSTRAINT "primary" PRIMARY KEY (id ASC),
FAMILY "primary" (id, first_name, last_name, score)
);