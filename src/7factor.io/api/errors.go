package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Validation error
type ValidationError struct {
	message string
}

func (e ValidationError) Error() string {
	return e.message
}

func NewValidationError(text string) error {
	return ValidationError{text}
}

func NewValidationErrorFromError(err error) error {
	return ValidationError{err.Error()}
}

// Private, package level functions for handling error messages being
// written in JSON to a writer buffer. Simple. I like this better than
// hacking some interface weirdness.
var errorTemplate = `{ "status": %v, "message": "%s" }`

func RenderErrorAndDeriveStatus(writer http.ResponseWriter, err error) {
	if err != nil {
		switch err.(type) {
		case ValidationError:
			RenderErrorWithStatus(writer, err, http.StatusBadRequest)
		default:
			RenderErrorWithStatus(writer, err, http.StatusInternalServerError)
		}
	}
}

func RenderErrorWithStatus(writer http.ResponseWriter, err error, status int) {
	writer.WriteHeader(status)
	buffer := new(bytes.Buffer)
	if err != nil {
		json.HTMLEscape(buffer, []byte(err.Error()))
	} else {
		json.HTMLEscape(buffer, []byte("no error specified"))
	}

	log.WithError(err).WithFields(log.Fields{
		"http_status_code": status,
		"message_body":     fmt.Sprintf(errorTemplate, status, buffer),
	}).Warn("Returning an error response code to client.")

	fmt.Fprintf(writer, errorTemplate, status, buffer)
}
