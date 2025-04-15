-- name: CreateSession :one
INSERT INTO user_sessions (
 id,
 user_uid,
 refresh_token,
 user_agent,
 client_ip,
 is_blocked,
 expired_at
) VALUES (
  $1, $2 ,$3,$4,$5,$6,$7
)
RETURNING *;

-- name: GetSessionById :one
SELECT * FROM user_sessions
WHERE id = $1 LIMIT 1;
