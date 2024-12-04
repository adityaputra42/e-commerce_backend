-- name: CreateUser :one
INSERT INTO users (
 uid,
 username,
 full_name,
 email,
 password,
 role,

) VALUES (
  $1, $2 ,$3 ,$4,$5,$6
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;
