package handlers

import "github.com/RaphaelParment/ppio-api/data"

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body data.GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body data.ValidationError
}

// A list of players
// swagger:response playersResponse
type playersResponseWrapper struct {
	// All current products
	// in: body
	Body []data.Player
}

// Data structure representing a single product
// swagger:response playerResponse
type playerResponseWrapper struct {
	// Newly created product
	// in: body
	Body data.Player
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
	Body data.Player
}

// swagger:parameters updatePlayer deletePlayer
type playerIDParamsWrapper struct {
	// The id of the player for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
