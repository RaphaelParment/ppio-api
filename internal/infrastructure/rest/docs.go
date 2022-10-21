package rest

//
//import (
//	"fmt"
//	"github.com/RaphaelParment/ppio-api/internal/domain/player/model"
//	"net/http"
//
//	"github.com/go-openapi/runtime/middleware"
//)
//
//// ErrInvalidProductPath is an error message when the product path is not valid
//var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")
//
//// GenericError is a generic error message returned by a server
//type GenericError struct {
//	Message string `json:"message"`
//}
//
//// ValidationError is a collection of validation error messages
//type ValidationError struct {
//	Messages []string `json:"messages"`
//}
//
//// Generic error message returned as a string
//// swagger:response errorResponse
//type errorResponseWrapper struct {
//	// Description of the error
//	// in: body
//	Body GenericError
//}
//
//// Validation errors defined as an array of strings
//// swagger:response errorValidation
//type errorValidationWrapper struct {
//	// Collection of the errors
//	// in: body
//	Body ValidationError
//}
//
//// A list of players
//// swagger:response playersResponse
//type playersResponseWrapper struct {
//	// All current products
//	// in: body
//	Body []model.Player
//}
//
//// Data structure representing a single product
//// swagger:response playerResponse
//type playerResponseWrapper struct {
//	// Newly created product
//	// in: body
//	Body model.Player
//}
//
//// No content is returned by this API endpoint
//// swagger:response noContentResponse
//type noContentResponseWrapper struct {
//}
//
//// swagger:parameters updatePlayer createPlayer
//type playerParamsWrapper struct {
//	// Player data structure to Update or Create.
//	// Note: the id field is ignored by update and create operations
//	// in: body
//	// required: true
//	Body model.Player
//}
//
//// swagger:parameters updatePlayer deletePlayer
//type playerIDParamsWrapper struct {
//	// The id of the player for which the operation relates
//	// in: path
//	// required: true
//	ID int `json:"id"`
//}
//
//func (s *server) handleRawDocsGet() http.Handler {
//	return http.FileServer(http.Dir("./"))
//}
//
//func (s *server) handleDocsGet() http.Handler {
//	opts := middleware.RedocOpts{SpecURL: "swagger.yaml"}
//	return middleware.Redoc(opts, nil)
//}
