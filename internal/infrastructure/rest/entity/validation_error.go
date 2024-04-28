package entity

import "github.com/RaphaelParment/ppio-api/internal/domain/match/validator"

type ValidationError struct {
	Problems map[string]string `json:"problems"`
}

func FromDomain(errors validator.ValidatorErrors) ValidationError {
	return ValidationError{
		Problems: errors.Problems(),
	}
}
