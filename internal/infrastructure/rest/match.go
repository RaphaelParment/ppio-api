package rest

import (
	"context"
	"encoding/json"
	"fmt"
	matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"
	playerModel "github.com/RaphaelParment/ppio-api/internal/domain/player/model"
	"github.com/RaphaelParment/ppio-api/internal/infrastructure/persistence/postgres/entity"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

func (s *server) HandleOneMatch() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestId, found := mux.Vars(r)["id"]
		if !found {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(requestId)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		match, err := s.matchService.HandleFindMatch(r.Context(), matchModel.Id(id))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		matchJSON := entity.MatchToJSON(match)
		m, err := json.Marshal(matchJSON)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		_, err = fmt.Fprintf(w, string(m))
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}
}

func (s *server) HandleMatches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			matches, err := s.handleGetMatches(r.Context())
			if err != nil {
				s.logger.Printf("failed to handle get matches; %s", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			m, err := json.Marshal(matches)
			if err != nil {
				s.logger.Printf("failed to marshal matches; %s", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(m)

		case http.MethodPost:
			match, err := s.handlePersistMatch(r.Context(), r.Body)
			if err != nil {
				s.logger.Printf("failed to handle persist match; %s", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			matchJson, err := json.Marshal(match)
			if err != nil {
				s.logger.Printf("failed to marshal match; %s", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(matchJson)

		default:
			log.Printf("unsupported method %q", r.Method)
			return
		}
	}
}

func (s *server) handleGetMatches(ctx context.Context) ([]entity.Match, error) {
	matches, err := s.matchService.HandleFindMatches(ctx)
	if err != nil {
		return nil, err
	}

	var matchesJson []entity.Match
	for _, match := range matches {
		matchesJson = append(matchesJson, entity.MatchToJSON(match))
	}

	return matchesJson, nil
}

func (s *server) handlePersistMatch(ctx context.Context, body io.ReadCloser) (entity.Match, error) {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return entity.Match{}, err
	}

	var inputMatch entity.Match

	err = json.Unmarshal(bodyBytes, &inputMatch)
	if err != nil {
		return entity.Match{}, err
	}

	s.logger.Println(inputMatch)

	match, err := s.matchService.HandlePersistMatch(
		ctx,
		playerModel.Id(inputMatch.PlayerOneId),
		playerModel.Id(inputMatch.PlayerTwoId),
		inputMatch.Datetime,
	)
	if err != nil {
		return entity.Match{}, err
	}

	matchJSON := entity.MatchToJSON(match)

	return matchJSON, nil
}
