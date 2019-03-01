package main

import (
	"github.com/husobee/vestigo"
	log "github.com/sirupsen/logrus"
	"goblog.com/api/db"
	"goblog.com/api/healthcheck"
	"net/http"
	"os"
	"time"
)

// Listen and go.
func main() {
	port := os.Getenv("PORT")
	log.Printf("We're about to be alive on port %s", port)

	router := vestigo.NewRouter()

	// Setting up router global CORS policy
	// These policy guidelines are overridable at a per resource level shown below
	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin:      []string{"*"},
		AllowCredentials: true,
		MaxAge:           3600 * time.Second,
		AllowHeaders:     []string{"Content-Type"},
	})

	// set up all of the database connections
	dbConn := db.GetDBOrPanic()

	// Please note that patterns for the URLs below must match
	// EXACTLY, including no trailing slashes.
	router.Get("/status", healthcheck.HealthCheck())
	router.Get("/dbaccess", healthcheck.DBAccess(dbConn))
	log.Fatal(http.ListenAndServe(":"+port, router))
}
