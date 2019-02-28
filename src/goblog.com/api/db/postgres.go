package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type Creds struct {
	host    string
	user    string
	pass    string
	port    string
	dbname  string
	options string
}

func NewDBConnection(creds *Creds) (*sql.DB, error) {
	connectionString := creds.ToConnectionString()
	return sql.Open("postgres", connectionString)
}

func GetCredsFromEnv() (*Creds, error) {
	user, ok := os.LookupEnv("DB_USER")
	if !ok {
		return nil, errors.New("no user specified in environment")
	}

	pass, ok := os.LookupEnv("DB_PASS")
	if !ok {
		return nil, errors.New("no password specified in environment")
	}

	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		return nil, errors.New("no host specified in environment")
	}

	port, ok := os.LookupEnv("DB_PORT")
	if !ok {
		return nil, errors.New("no port specified in environment")
	}

	dbname, ok := os.LookupEnv("DB_NAME")
	if !ok {
		return nil, errors.New("no database specified in environment")
	}

	options := os.Getenv("DB_OPTN")

	return &Creds{host, user, pass, port, dbname, options}, nil
}

const dbConnectionTemplate = `postgres://%s:%s@%s:%s/%s`

func (c Creds) ToConnectionString() string {
	connStr := fmt.Sprintf(dbConnectionTemplate, c.user, c.pass, c.host, c.port, c.dbname)

	connStr = appendOptions(connStr, c.options)

	return connStr
}

func appendOptions(connStr string, options string) string {
	if len(options) > 0 {
		connStr += "?" + options
	}
	return connStr
}
