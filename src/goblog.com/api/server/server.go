package main

import (
	"goblog.com/api"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/husobee/vestigo"
)

// Listen and go.
func main() {
	port := os.Getenv("PORT")
	log.Printf("We're about to be alive on port %s", port)

	router := vestigo.NewRouter()

	// Please note that patterns for the URLs below must match
	// EXACTLY, including no trailing slashes.
	router.Get("/status", api.HealthCheck())
	log.Fatal(http.ListenAndServe(":"+port, router))
}
