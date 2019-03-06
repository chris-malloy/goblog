package models

const insertUserSQL = `
    INSERT INTO users (first_name, last_name, email, encrypted_password, sign_in_count, last_sign_in_at, created_time_stamp, updated_time_stamp)
    VALUES ($1, $2, $3, $4, 0, null, $5, null);
`

const selectUserSQL = `
	SELECT id, email, first_name, last_name, last_sign_in_at, sign_in_count
	FROM users
    WHERE email = $1;
`
