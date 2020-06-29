package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/IdleTradingHeroServer/constants"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

const (
	maxBytesInRequest = 1048576
)

type errorPayload struct {
	Status int
	Errors []*errorResponse
}

type errorResponse struct {
	Code    string
	Message string
}

func (er *errorResponse) Error() string {
	return er.Message
}

func (er *errorPayload) Error() string {
	message := ""

	for _, err := range er.Errors {
		message += fmt.Sprintln(err.Error())
	}

	return message
}

func respondWithErrorPayloadFromString(logger *log.Logger, w http.ResponseWriter, status int, code string, errMsg string) {
	payload := generateErrorPayloadFromString(status, code, errMsg)

	respondWithErrorPayload(logger, w, payload)
}

func respondWithErrorPayloadFromErrorResponse(logger *log.Logger, w http.ResponseWriter, status int, err *errorResponse) {
	errorResponses := []*errorResponse{err}
	respondWithErrorPayloadFromErrorResponses(logger, w, status, errorResponses)
}

func respondWithErrorPayloadFromErrorResponses(logger *log.Logger, w http.ResponseWriter, status int, errorResponses []*errorResponse) {
	payload := generateErrorPayloadFromErrorResponses(status, errorResponses)

	respondWithErrorPayload(logger, w, payload)
}

func respondWithErrorPayloadFromErrors(logger *log.Logger, w http.ResponseWriter, status int, code string, errors []error) {
	payload := generateErrorPayloadFromErrors(status, code, errors)

	respondWithErrorPayload(logger, w, payload)
}

func respondWithErrorPayload(logger *log.Logger, w http.ResponseWriter, payload *errorPayload) {
	jsonPayload, err := serialiseErrorPayload(payload)

	if err != nil {
		logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(payload.Status)
	w.Write(jsonPayload)
}

func serialiseErrorPayloadFromString(status int, code, errMsg string) ([]byte, error) {
	payload := generateErrorPayloadFromString(status, code, errMsg)

	return serialiseErrorPayload(payload)
}

func serialiseErrorPayloadFromErrors(status int, errResponses []*errorResponse) ([]byte, error) {
	payload := &errorPayload{
		Status: status,
		Errors: errResponses,
	}

	return serialiseErrorPayload(payload)
}

func serialiseErrorPayload(payload *errorPayload) ([]byte, error) {
	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	return []byte(jsonPayload), nil
}

func generateErrorPayloadFromString(status int, code, errMsg string) *errorPayload {
	err := &errorResponse{
		Code:    code,
		Message: errMsg,
	}
	errors := []*errorResponse{err}

	return generateErrorPayloadFromErrorResponses(status, errors)
}

func generateErrorPayloadFromErrors(status int, code string, errors []error) *errorPayload {
	errorResponses := make([]*errorResponse, len(errors))

	for i, err := range errors {
		errorResponses[i] = &errorResponse{
			Code:    code,
			Message: err.Error(),
		}
	}

	return generateErrorPayloadFromErrorResponses(status, errorResponses)
}

func generateErrorPayloadFromErrorResponses(status int, errorResponses []*errorResponse) *errorPayload {
	payload := &errorPayload{
		Status: status,
		Errors: errorResponses,
	}

	return payload
}

func respondJSON(w http.ResponseWriter, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) *errorPayload {
	if r.Header.Get("Content-Type") != "" {
		value := r.Header.Get("Content-Type")
		if !strings.Contains(value, "application/json") {
			errMsg := "Content-Type header is not application/json"
			return generateErrorPayloadFromString(http.StatusBadRequest, constants.ErrorCodeInvalidContentType, errMsg)
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxBytesInRequest)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			errMsg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return generateErrorPayloadFromString(http.StatusBadRequest, constants.ErrorCodeMalformedJSON, errMsg)

		case errors.Is(err, io.ErrUnexpectedEOF):
			errMsg := fmt.Sprintf("Request body contains badly-formed JSON")
			return generateErrorPayloadFromString(http.StatusBadRequest, constants.ErrorCodeMalformedJSON, errMsg)

		case errors.As(err, &unmarshalTypeError):
			errMsg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return generateErrorPayloadFromString(http.StatusBadRequest, constants.ErrorCodeMalformedJSON, errMsg)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			errMsg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return generateErrorPayloadFromString(http.StatusBadRequest, constants.ErrorCodeMalformedJSON, errMsg)

		case errors.Is(err, io.EOF):
			errMsg := "Request body must not be empty"
			return generateErrorPayloadFromString(http.StatusBadRequest, constants.ErrorCodeMalformedJSON, errMsg)

		case err.Error() == "http: request body too large":
			errMsg := "Request body must not be larger than 1MB"
			return generateErrorPayloadFromString(http.StatusRequestEntityTooLarge, constants.ErrorCodeRequestTooLarge, errMsg)

		default:
			panic(err)
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		errMsg := "Request body must only contain a single JSON object"
		return generateErrorPayloadFromString(http.StatusBadRequest, constants.ErrorCodeMalformedJSON, errMsg)
	}

	return nil
}

func getUserIDFromJWT(r *http.Request) string {
	j := r.Context().Value("user")
	accessToken := j.(*jwt.Token)
	userID := accessToken.Claims.(jwt.MapClaims)["userId"].(string)

	return userID
}

func validateStruct(s interface{}, w http.ResponseWriter, logger *log.Logger) bool {
	// check request
	if errs := validate.Struct(s); errs != nil {
		validationErrs, _ := errs.(validator.ValidationErrors)

		errorResponses := make([]*errorResponse, len(validationErrs))

		for i, validationErr := range validationErrs {
			errorResponses[i] = &errorResponse{
				Code:    constants.ErrorCodeInvalidField,
				Message: validationErr.Translate(enTranslator),
			}
		}

		respondWithErrorPayloadFromErrorResponses(logger, w, http.StatusUnprocessableEntity, errorResponses)
		return false
	}

	return true
}
