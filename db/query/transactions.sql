-- name: CreateTransaction :one
INSERT INTO transactions (
 tx_id,
 address_id,
 shipping_id,
 shipping_price,
 status
) VALUES (
  $1, $2 ,$3, $4, $5
)
RETURNING *;


-- name: GetTransaction :one
SELECT * FROM transactions
WHERE tx_id = $1 LIMIT 1;

-- name: GetTransactionForUpdate :one
SELECT * FROM transactions
WHERE tx_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTransaction :many
SELECT * FROM transactions
ORDER BY tx_id
LIMIT $1
OFFSET $2;

-- name: UpdateTransaction :one
UPDATE transactions
 set status = $2
WHERE tx_id = $1
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE tx_id = $1;