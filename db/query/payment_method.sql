-- name: CreatePaymentMethod :one
INSERT INTO payment_method (
 account_name,
 account_number,
 bank_name,
 bank_images
) VALUES (
  $1, $2 ,$3, $4
)
RETURNING *;

-- name: GetPaymentMethod :one
SELECT * FROM payment_method
WHERE deleted_at IS NULL AND id = $1 LIMIT 1;

-- name: GetPaymentMethodForUpdate :one
SELECT * FROM payment_method
WHERE deleted_at IS NULL AND id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListPaymentMethod :many
SELECT * FROM payment_method
WHERE deleted_at IS NULL
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePaymentMethod :one
UPDATE payment_method
 set account_name = $2,
 account_number = $3,
 bank_name = $4,
 bank_images = $5
WHERE id = $1
RETURNING *;

-- name: DeletePaymentMethod :exec
UPDATE payment_method
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;