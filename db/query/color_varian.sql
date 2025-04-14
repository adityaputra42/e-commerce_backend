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
SELECT 
  cv.id AS id,
  cv.product_id AS product_id,
  cv.name AS name,
  cv.color AS color,
  cv.images AS images,
  cv.updated_at AS updated_at,
  cv.created_at AS created_at,
  jsonb_agg(
      jsonb_build_object(
          'id', sv.id,
          'color_varian_id', sv.color_varian_id,
          'size', sv.size,
          'stock', sv.stock,
          'updated_at', sv.updated_at,
          'created_at', sv.created_at
           )
     )AS size_varians
 FROM color_varians cv
LEFT JOIN 
    size_varians sv ON cv.id = sv.color_varian_id AND sv.deleted_at IS NULL
WHERE cv.deleted_at IS NULL AND cv.id = $1 LIMIT 1;

-- name: GetColorVarianProductForUpdate :one
SELECT * FROM color_varians
WHERE deleted_at IS NULL AND id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListColorVarianProduct :many
SELECT 
  cv.id AS id,
  cv.product_id AS product_id,
  cv.name AS name,
  cv.color AS color,
  cv.images AS images,
  cv.updated_at AS updated_at,
  cv.created_at AS created_at,
  jsonb_agg(
      jsonb_build_object(
          'id', sv.id,
          'color_varian_id', sv.color_varian_id,
          'size', sv.size,
          'stock', sv.stock,
          'updated_at', sv.updated_at,
          'created_at', sv.created_at
           )
     )AS size_varians

FROM color_varians cv 
LEFT JOIN 
    size_varians sv ON cv.id = sv.color_varian_id AND sv.deleted_at IS NULL
WHERE cv.deleted_at IS NULL AND cv.product_id = $1
ORDER BY cv.id
LIMIT $2
OFFSET $3 ;

-- name: UpdateColorVarianProduct :one
UPDATE color_varians
set name = $2,
color = $3,
images = $4
WHERE id = $1
RETURNING *;

-- name: DeleteColorVarianProduct :exec
UPDATE color_varians
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;