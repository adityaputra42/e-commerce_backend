-- name: CreatePayment :one
INSERT INTO payment (
 payment_method_id,
 transaction_id,
 total_payment,
 status
) VALUES (
  $1, $2 ,$3, $4
)
RETURNING *;


-- name: GetPayment :one
SELECT * FROM payment
WHERE id = $1 LIMIT 1;

-- name: GetPaymentForUpdate :one
SELECT * FROM payment
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListPayment :many
SELECT * FROM payment
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePayment :one
UPDATE payment
 set status = $2
WHERE id = $1
RETURNING *;

-- name: DeletePayment :exec
DELETE FROM payment
WHERE id = $1;