-- name: CreateOrder :one
INSERT INTO orders (
 id,
 transaction_id,
 product_id,
 color_varian_id,
 size_varian_id,
 unit_price,
 subtotal,
 quantity,
 status
) VALUES (
  $1, $2 ,$3, $4, $5, $6, $7, $8, $9
)
RETURNING *;


-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: GetOrderForUpdate :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListOrder :many
SELECT * FROM orders
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateOrder :one
UPDATE orders
 set status = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;