DROP DATABASE IF EXISTS ppio CASCADE;

CREATE DATABASE ppio;

USE ppio;

CREATE TABLE player (
    id serial PRIMARY KEY,
    first_name VARCHAR(32) NOT NULL,
    last_name VARCHAR(32) NOT NULL,
    email VARCHAR (64) UNIQUE NOT NULL,
    points INT
);

CREATE TABLE validation (
    id INTEGER NOT NULL,
    description STRING NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    FAMILY "primary" (id, description)
);

INSERT INTO validation(id, description) VALUES (0, 'in process');
INSERT INTO validation(id, description) VALUES (1, 'validated');

CREATE TABLE game (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    player1_id UUID NOT NULL,
    player2_id UUID NOT NULL,
    winner_id UUID NOT NULL,
    datetime TIMESTAMP WITH TIME ZONE NULL,
    validation_state INTEGER NOT NULL DEFAULT 0,
    edited_by_id UUID NOT NULL,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    CONSTRAINT player1_fk FOREIGN KEY (player1_id) REFERENCES player (id),
    INDEX game_player1_id_idx (player1_id ASC),
    CONSTRAINT player2_fk FOREIGN KEY (player2_id) REFERENCES player (id),
    INDEX game_player2_id_idx (player2_id ASC),
    CONSTRAINT winner_fk FOREIGN KEY (winner_id) REFERENCES player (id),
    INDEX game_winner_id_idz (winner_id ASC),
    CONSTRAINT game_validation_state_fk FOREIGN KEY (validation_state) REFERENCES validation (id),
    INDEX game_validation_state_idz (validation_state),
    CONSTRAINT edited_by_id_fk FOREIGN KEY (edited_by_id) REFERENCES player (id),
    INDEX game_edited_by_id_idz (edited_by_id ASC),
    FAMILY "primary" (id, player1_id, player2_id, datetime)
);

CREATE TABLE set (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    game_id UUID NOT NULL,
    score1 INT NOT NULL DEFAULT 0,
    score2 INT NOT NULL DEFAULT 0,
    CONSTRAINT "primary" PRIMARY KEY (id ASC),
    CONSTRAINT game_id_fk FOREIGN KEY (game_id) REFERENCES game(id) ON DELETE CASCADE,
    FAMILY "primary" (id, game_id, score1, score2)
);

CREATE USER IF NOT EXISTS ppio_user;
GRANT ALL ON ppio.* TO ppio_user;
