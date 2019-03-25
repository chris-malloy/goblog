package models

const insertUserSQL = `
    INSERT INTO users (first_name, last_name, email, encrypted_password, sign_in_count, last_sign_in_at, created_time_stamp, updated_time_stamp)
    VALUES ($1, $2, $3, $4, 0, null, $5, null);
`

const selectUserByIdSQL = `
	SELECT id, email, first_name, last_name, last_sign_in_at, sign_in_count
	FROM users 
	WHERE id = $1;
`

const selectUserByEmailSQL = `
	SELECT id, email, first_name, last_name, last_sign_in_at, sign_in_count
	FROM users
    WHERE email = $1;
`

const updateUserByIdSQL = `
	UPDATE users set email = $2, first_name = $3, last_name = $4, updated_time_stamp = NOW()
	WHERE id = $1;
`

const deleteUserByIdSQL = `
	DELETE FROM users
	WHERE id = $1;
`

const getEncryptedPasswordSQL = `
	SELECT encrypted_password
	FROM users
	WHERE email = $1;
`
