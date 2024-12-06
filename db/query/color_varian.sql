-- name: CreateColorVarianProduct :one
INSERT INTO color_varians (
 product_id,
 name,
 color,
 images
) VALUES (
  $1, $2 ,$3 ,$4
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
set name = $2,
color = $3,
images = $4
WHERE id = $1
RETURNING *;

-- name: DeleteColorVarianProduct :exec
DELETE FROM color_varians
WHERE id = $1;