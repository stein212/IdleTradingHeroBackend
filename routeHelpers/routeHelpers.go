package routehelpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/IdleTradingHeroServer/constants"
)

const (
	maxBytesInRequest = 1048576
)

type ErrorPayload struct {
	Status int
	Errors []*ErrorResponse
}

type ErrorResponse struct {
	Code    string
	Message string
}

func (er *ErrorResponse) Error() string {
	return er.Message
}

func (er *ErrorPayload) Error() string {
	message := ""

	for _, err := range er.Errors {
		message += fmt.Sprintln(err.Error())
	}

	return message
}

func RespondWithErrorPayloadFromString(logger *log.Logger, w http.ResponseWriter, status int, code string, errMsg string) {
	payload := generateErrorPayloadFromString(status, code, errMsg)

	RespondWithErrorPayload(logger, w, payload)
}

func RespondWithErrorPayloadFromErrorResponse(logger *log.Logger, w http.ResponseWriter, status int, err *ErrorResponse) {
	errorResponses := []*ErrorResponse{err}
	RespondWithErrorPayloadFromErrorResponses(logger, w, status, errorResponses)
}

func RespondWithErrorPayloadFromErrorResponses(logger *log.Logger, w http.ResponseWriter, status int, errorResponses []*ErrorResponse) {
	payload := generateErrorPayloadFromErrorResponses(status, errorResponses)

	RespondWithErrorPayload(logger, w, payload)
}

func RespondWithErrorPayloadFromErrors(logger *log.Logger, w http.ResponseWriter, status int, code string, errors []error) {
	payload := generateErrorPayloadFromErrors(status, code, errors)

	RespondWithErrorPayload(logger, w, payload)
}

func RespondWithErrorPayload(logger *log.Logger, w http.ResponseWriter, payload *ErrorPayload) {
	jsonPayload, err := SerialiseErrorPayload(payload)

	if err != nil {
		logger.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(payload.Status)
	w.Write(jsonPayload)
}

func SerialiseErrorPayloadFromString(status int, code, errMsg string) ([]byte, error) {
	payload := generateErrorPayloadFromString(status, code, errMsg)

	return SerialiseErrorPayload(payload)
}

func SerialiseErrorPayloadFromErrors(status int, errResponses []*ErrorResponse) ([]byte, error) {
	payload := &ErrorPayload{
		Status: status,
		Errors: errResponses,
	}

	return SerialiseErrorPayload(payload)
}

func SerialiseErrorPayload(payload *ErrorPayload) ([]byte, error) {
	jsonPayload, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	return []byte(jsonPayload), nil
}

func generateErrorPayloadFromString(status int, code, errMsg string) *ErrorPayload {
	err := &ErrorResponse{
		Code:    code,
		Message: errMsg,
	}
	errors := []*ErrorResponse{err}

	return generateErrorPayloadFromErrorResponses(status, errors)
}

func generateErrorPayloadFromErrors(status int, code string, errors []error) *ErrorPayload {
	errorResponses := make([]*ErrorResponse, len(errors))

	for i, err := range errors {
		errorResponses[i] = &ErrorResponse{
			Code:    code,
			Message: err.Error(),
		}
	}

	return generateErrorPayloadFromErrorResponses(status, errorResponses)
}

func generateErrorPayloadFromErrorResponses(status int, errorResponses []*ErrorResponse) *ErrorPayload {
	payload := &ErrorPayload{
		Status: status,
		Errors: errorResponses,
	}

	return payload
}

func RespondJSON(w http.ResponseWriter, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payload)
}

func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) *ErrorPayload {
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
