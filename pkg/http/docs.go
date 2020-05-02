package http

import (
	"net/http"

	"github.com/RaphaelParment/ppio-api/pkg/core"
	"github.com/go-openapi/runtime/middleware"
)

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body core.GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body core.ValidationError
}

// A list of players
// swagger:response playersResponse
type playersResponseWrapper struct {
	// All current products
	// in: body
	Body []core.Player
}

// Data structure representing a single product
// swagger:response playerResponse
type playerResponseWrapper struct {
	// Newly created product
	// in: body
	Body core.Player
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updatePlayer createPlayer
type playerParamsWrapper struct {
	// Player data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body core.Player
}

// swagger:parameters updatePlayer deletePlayer
type playerIDParamsWrapper struct {
	// The id of the player for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

func (s *server) handleRawDocsGet() http.Handler {
	return http.FileServer(http.Dir("./"))
}

func (s *server) handleDocsGet() http.Handler {
	opts := middleware.RedocOpts{SpecURL: "swagger.yaml"}
	return middleware.Redoc(opts, nil)
}
