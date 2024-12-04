-- name: CreateShipping :one
INSERT INTO shippings (
 name,
 price,
 state
) VALUES (
  $1, $2 ,$3
)
RETURNING *;

-- name: GetShipping :one
SELECT * FROM shippings
WHERE id = $1 LIMIT 1;

-- name: GetShippingForUpdate :one
SELECT * FROM shippings
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListShipping :many
SELECT * FROM shippings
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateShipping :one
UPDATE shippings
 set name = $2,
 price = $3,
 state = $4
WHERE id = $1
RETURNING *;

-- name: DeleteShipping :exec
DELETE FROM shippings
WHERE id = $1;