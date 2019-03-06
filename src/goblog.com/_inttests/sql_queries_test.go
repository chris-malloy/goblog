package _inttests

import (
	"database/sql"
	. "github.com/onsi/ginkgo"
)

const insertTestUserSQL = `
	INSERT INTO users (first_name, last_name, email, encrypted_password, sign_in_count, created_time_stamp, updated_time_stamp) 
	VALUES ('Chris', 'Malloy', 'christopher.malloy@7factor.io', '$2a$10$2HAOhibpvDn/iinGWT4ME.n0Upe45b3.5hHRG9FtkyfICoM9ilzta', 0, NOW(), NULL)
	RETURNING id;
`

func insertTestUser(db *sql.DB) int64 {
	var id int64

	err := db.QueryRow(insertTestUserSQL).Scan(&id)
	if err != nil {
		GinkgoT().Fatalf("Caught error while attempting to insert test user: %v", err.Error())
	}

	return id
}
