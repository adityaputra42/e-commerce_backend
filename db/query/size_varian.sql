-- name: CreateSizeVarianProduct :one
INSERT INTO size_varians (
 color_varian_id,
 size,
 stock
) VALUES (
  $1, $2 ,$3
)
RETURNING *;


-- name: GetSizeVarianProduct :one
SELECT * FROM size_varians
WHERE id = $1 LIMIT 1;

-- name: GetSizeVarianProductForUpdate :one
SELECT * FROM size_varians
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListSizeVarianProduct :many
SELECT * FROM size_varians
WHERE color_varian_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateSizeVarianProduct :one
UPDATE size_varians
 set size = $2,
 stock = $3
WHERE id = $1
RETURNING *;

-- name: DeleteSizeVarianProduct :exec
DELETE FROM size_varians
WHERE id = $1;