-- name: CreateProduct :one
INSERT INTO products (
 category_id,
 name,
 description,
 images,
 rating,
 price
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;


-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: GetProductForUpdate :one
SELECT * FROM products
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListProduct :many
SELECT * FROM products
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProduct :one
UPDATE products
 set category_id = $2,
 name = $3,
 description= $4,
 images= $5,
 rating= $6,
 price= $7
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = $1;