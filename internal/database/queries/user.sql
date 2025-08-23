-- name: CreateUser :one
INSERT INTO users(email, password, first_name, last_name, dob, role)
VALUES(@email, @password, @first_name, @last_name, @dob, @role)
RETURNING id;

-- name: GetUserByEmail :one
SELECT id, password
FROM users
WHERE email = @email;
