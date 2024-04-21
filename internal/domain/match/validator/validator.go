package validator

import matchModel "github.com/RaphaelParment/ppio-api/internal/domain/match/model"

type matchValidator struct {
}

func NewMatchValidator() *matchValidator {
	return &matchValidator{}
}

func (v *matchValidator) ValidateMatch(match matchModel.Match) map[string]string {
	return nil
}
