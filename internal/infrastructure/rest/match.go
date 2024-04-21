package rest

import (
	"encoding/json"
	"errors"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	restEntity "github.com/RaphaelParment/ppio-api/internal/infrastructure/rest/entity"
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

	err = c.JSON(http.StatusOK, restEntity.MatchToJSON(match))
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

	var matchesJSON []restEntity.Match
	for _, match := range matches {
		matchesJSON = append(matchesJSON, restEntity.MatchToJSON(match))
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

	var inputMatch restEntity.Match
	err = json.Unmarshal(bodyBytes, &inputMatch)
	if err != nil {
		s.logger.Printf("failed to unmarshal body into add match; %s", err)
		return err
	}

	match, err := restEntity.MatchFromJSON(inputMatch)
	if err != nil {
		s.logger.Printf("failed to convert to domain match; %s", err)
		return err
	}

	id, err := s.matchService.HandlePersistMatch(c.Request().Context(), match)
	if err != nil {
		s.logger.Printf("failed to persist match, result and scores; %s", err)
		return err
	}

	err = c.JSON(http.StatusOK, id)
	if err != nil {
		s.logger.Printf("failed to marshal match id; %s", err)
		return err
	}

	return nil
}

func (s *server) HandleUpdateOneMatch(c echo.Context) error {
	matchId := c.Param("id")
	if matchId == "" {
		s.logger.Printf("missing match id")
		return errors.New("missing match id")
	}

	id, err := strconv.Atoi(matchId)
	if err != nil {
		s.logger.Printf("failed to convert match id to int; %s", err)
		return err
	}

	bodyBytes, err := io.ReadAll(c.Request().Body)
	if err != nil {
		s.logger.Printf("failed to read body; %s", err)
		return err
	}

	var inputPatchedMatch restEntity.MatchPatch
	err = json.Unmarshal(bodyBytes, &inputPatchedMatch)
	if err != nil {
		s.logger.Printf("failed to unmarshal body into match patch; %s", err)
		return err
	}

	patchedMatch, err := restEntity.MatchPatchFromJSON(inputPatchedMatch)
	if err != nil {
		s.logger.Printf("failed to convert match patch to domain match patch; err: %s", err)
		return err
	}

	match, err := s.matchService.HandleUpdateOneMatch(c.Request().Context(), matchModel.Id(id), patchedMatch)
	if err != nil {
		s.logger.Printf("failed to update match id %d, err: %s", id, err)
		return err
	}

	err = c.JSON(http.StatusOK, restEntity.MatchToJSON(match))
	if err != nil {
		s.logger.Printf("failed to return json match response; %s", err)
		return err
	}

	return nil
}
