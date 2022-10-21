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

func (s *server) HandleGetOneMatch(c echo.Context) error {
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

	match, err := s.matchService.HandleFindMatch(c.Request().Context(), matchModel.Id(id))
	if err != nil {
		s.logger.Printf("failed to find match; %s", err)
		return err
	}

	err = c.JSON(http.StatusOK, entity.MatchToJSON(match))
	if err != nil {
		s.logger.Printf("failed to return json match response; %s", err)
		return err
	}

	return nil
}

func (s *server) HandleGetAllMatches(c echo.Context) error {
	matches, err := s.matchService.HandleFindMatches(c.Request().Context())
	if err != nil {
		s.logger.Printf("failed to find all matches; %s", err)
		return err
	}

	var matchesJSON []entity.Match
	for _, match := range matches {
		matchesJSON = append(matchesJSON, entity.MatchToJSON(match))
	}

	err = c.JSON(http.StatusOK, matchesJSON)
	if err != nil {
		s.logger.Printf("failed to marshal matches; %s", err)
		return err
	}

	return nil
}

func (s *server) HandleAddOneMatch(c echo.Context) error {
	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		s.logger.Printf("failed to read body; %s", err)
		return err
	}

	var inputMatch entity.Match
	err = json.Unmarshal(bodyBytes, &inputMatch)
	if err != nil {
		s.logger.Printf("failed to unmarshal body into match; %s", err)
		return err
	}

	match, err := s.matchService.HandlePersistMatch(
		c.Request().Context(),
		playerModel.Id(inputMatch.PlayerOneId),
		playerModel.Id(inputMatch.PlayerTwoId),
		inputMatch.Datetime,
	)
	if err != nil {
		s.logger.Printf("failed to persist match; %s", err)
		return err
	}

	err = c.JSON(http.StatusOK, entity.MatchToJSON(match))
	if err != nil {
		s.logger.Printf("failed to marshal match; %s", err)
		return err
	}

	return nil
}
