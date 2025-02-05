-- name: CreateUser :one
INSERT INTO users (
    id, username, password
) 
VALUES ($1, $2, $3)
RETURNING *;

-- name: DeleteUserById :one
DELETE FROM users
WHERE id = $1
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

