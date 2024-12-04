-- name: CreateUser :one
INSERT INTO address (
 id,
 uid,
 full_name,
 email,
 password,
 role,

) VALUES (
  $1, $2 ,$3 ,$4,$5,$6
)
RETURNING *;
