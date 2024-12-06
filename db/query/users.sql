-- name: CreateUser :one
INSERT INTO users (
 uid,
 username,
 full_name,
 email,
 password,
 role
) VALUES (
  $1, $2 ,$3 ,$4,$5,$6
)
RETURNING *;

-- name: GetUserLogin :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE uid = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListUser :many
SELECT * FROM users
WHERE role = $1
ORDER BY uid 
LIMIT $2
OFFSET $3;

-- name: UpdateUser :one
UPDATE users
 set password = $2
WHERE uid = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE uid = $1;