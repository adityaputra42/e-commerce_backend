-- name: CreatePayment :one
INSERT INTO payment (
 transaction_id,
 total_payment,
 status
) VALUES (
  $1, $2 ,$3
)
RETURNING *;


-- name: GetPayment :one
SELECT * FROM payment
WHERE deleted_at IS NOT NULL AND id = $1 LIMIT 1;

-- name: GetPaymentForUpdate :one
SELECT * FROM payment
WHERE deleted_at IS NOT NULL AND id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListPayment :many
SELECT * FROM payment
WHERE deleted_at IS NOT NULL
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePayment :one
UPDATE payment
 set status = $2
WHERE id = $1
RETURNING *;

-- name: DeletePayment :exec
UPDATE payment
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;