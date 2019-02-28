package db

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

func GetDBOrPanic() *sql.DB {
	creds, err := GetCredsFromEnv()
	if err != nil {
		log.WithError(err).Fatal("Unable to start server. DBCreds are invalid.")
	}

	connection, err := NewDBConnection(creds)
	if err != nil {
		log.WithError(err).Fatal("Unable to start server. Cannot create a DB instance.")
	}

	return connection
}
