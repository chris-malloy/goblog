package handlers

import (
	"encoding/json"
	"net/http"
)

// StatusMessage struct for representing a status message
type StatusMessage struct {
	Gtg bool   `json:"gtg"`
	Err string `json:"err"`
}

// HealthCheck returns a message signaling if the service can take
// traffic or not. This does not tell you if the service is healthy or not.
func HealthCheck() http.HandlerFunc {
	// Type coercion black magic.
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(StatusMessage{true, "none"})
	}
}
