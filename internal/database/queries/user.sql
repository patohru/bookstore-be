-- name: CreateUser :one
INSERT INTO users(email, password, first_name, last_name, dob)
VALUES(@email, @password, @first_name, @last_name, @dob)
RETURNING id;

-- name: GetUserByEmail :one
SELECT id, password
FROM users
WHERE email = @email;
