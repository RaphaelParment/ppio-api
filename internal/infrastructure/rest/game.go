package rest

import (
	"encoding/json"
	"errors"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence/postgres/entity"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strconv"
)

func (s *server) HandleGetOneGame(c echo.Context) error {
	requestId := c.Param("id")
	if requestId == "" {
		s.logger.Printf("missing id")
		return errors.New("missing id")
	}

	id, err := strconv.Atoi(requestId)
	if err != nil {
		s.logger.Printf("failed to convert request id to int; %s", err)
		return err
	}

	match, err := s.gameService.HandleFindGame(c.Request().Context(), matchModel.Id(id))
	if err != nil {
		s.logger.Printf("failed to find match; %s", err)
		return err
	}

	err = c.JSON(http.StatusOK, entity.GameToJSON(match))
	if err != nil {
		s.logger.Printf("failed to return json game response; %s", err)
		return err
	}

	return nil
}

func (s *server) HandleGetAllGames(c echo.Context) error {
	games, err := s.gameService.HandleFindGames(c.Request().Context())
	if err != nil {
		s.logger.Printf("failed to find all games; %s", err)
		return err
	}

	var gamesJson []entity.Match
	for _, match := range games {
		gamesJson = append(gamesJson, entity.GameToJSON(match))
	}

	err = c.JSON(http.StatusOK, gamesJson)
	if err != nil {
		s.logger.Printf("failed to marshal games; %s", err)
		return err
	}

	return nil
}

func (s *server) HandleAddOneGame(c echo.Context) error {
	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		s.logger.Printf("failed to read body; %s", err)
		return err
	}

	var inputGame entity.Match
	err = json.Unmarshal(bodyBytes, &inputGame)
	if err != nil {
		s.logger.Printf("failed to unmarshal body into game; %s", err)
		return err
	}

	game, err := s.gameService.HandlePersistGame(
		c.Request().Context(),
		playerModel.Id(inputGame.PlayerOneId),
		playerModel.Id(inputGame.PlayerTwoId),
		inputGame.Datetime,
	)
	if err != nil {
		s.logger.Printf("failed to persist game; %s", err)
		return err
	}

	err = c.JSON(http.StatusOK, entity.GameToJSON(game))
	if err != nil {
		s.logger.Printf("failed to marshal game; %s", err)
		return err
	}

	return nil
}
