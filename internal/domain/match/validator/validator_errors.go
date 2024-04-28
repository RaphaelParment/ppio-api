package validator

import (
	"fmt"
	"strings"
)

type ValidatorErrors struct {
	problems map[string]string
}

func NewValidatorError(problems map[string]string) ValidatorErrors {
	return ValidatorErrors{problems: problems}
}

func (ve ValidatorErrors) Problems() map[string]string {
	return ve.problems
}

func (ve ValidatorErrors) Error() string {
	var errors []string
	for errorKey, errorMsg := range ve.problems {
		errors = append(errors, fmt.Sprintf("%s: %s", errorKey, errorMsg))
	}

	return strings.Join(errors, ";")
}
