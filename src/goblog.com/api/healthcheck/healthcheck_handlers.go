package healthcheck

import (
	"encoding/json"
	"net/http"
)

type StatusMessage struct {
	Gtg bool   `json:"gtg"`
	Err string `json:"err"`
}

func HealthCheck() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(StatusMessage{true, "none"})
	}
}
