-- name: CreateColorVarianProduct :one
INSERT INTO color_varians (
 product_id,
 color,
 images
) VALUES (
  $1, $2 ,$3
)
RETURNING *;


-- name: GetColorVarianProduct :one
SELECT * FROM color_varians
WHERE id = $1 LIMIT 1;

-- name: GetColorVarianProductForUpdate :one
SELECT * FROM color_varians
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListColorVarianProduct :many
SELECT * FROM color_varians
WHERE product_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateColorVarianProduct :one
UPDATE color_varians
 set color = $2,
 images = $3
WHERE id = $1
RETURNING *;

-- name: DeleteColorVarianProduct :exec
DELETE FROM color_varians
WHERE id = $1;