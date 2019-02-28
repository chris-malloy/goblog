package healthcheck

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type StatusMessage struct {
	Up  bool   `json:"up"`
	Err string `json:"err"`
}

func HealthCheck() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(StatusMessage{true, "none"})
	}
}

func DBAccess(db *sql.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		err := db.Ping()
		if err != nil {
			writer.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(writer).Encode(StatusMessage{false, err.Error()})
		} else {
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(StatusMessage{true, "none"})
		}
	}
}
