-- name: CreateCategories :one
INSERT INTO categories (
 name,
 icon
) VALUES (
  $1, $2 
)
RETURNING *;


-- name: GetCategories :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: GetCategoriesForUpdate :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListCategories :many
SELECT * FROM categories
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCategories :one
UPDATE categories
 set name = $2,
 icon = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCategories :exec
DELETE FROM categories
WHERE id = $1;