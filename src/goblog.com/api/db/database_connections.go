package db

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"goblog.com/api/models"
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

func GetAuthorizerOrPanic(db *sql.DB) models.Authorizer {
	authorizer, err := models.NewAuthorizer(db)
	if err != nil {
		log.WithError(err).Fatal("Unable to start server. Cannot create authorization module.")
	}

	return authorizer
}

func GetUserManagerOrPanic(db *sql.DB) models.UserCRUD {
	userManger, err := models.NewUserManager(db)
	if err != nil {
		log.WithError(err).Fatal("Unable to start server. Cannot create user manager.")
	}

	return userManger
}
