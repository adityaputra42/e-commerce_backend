-- name: CreateAddress :one
INSERT INTO address (
 uid,
 recipient_name,
 recipient_phone_number,
 province,
 city,
 district,
 village,
 postal_code,
 full_address
) VALUES (
  $1, $2 ,$3 ,$4,$5,$6,$7,$8,$9
)
RETURNING *;


-- name: GetAddress :one
SELECT * FROM address
WHERE id = $1 LIMIT 1;

-- name: GetAddressForUpdate :one
SELECT * FROM address
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAddress :many
SELECT * FROM address
WHERE uid = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAddress :one
UPDATE address
 set recipient_name = $2,
 recipient_phone_number = $3,
 province = $4,
 city = $5,
 district = $6,
 village = $7,
 postal_code = $8,
 full_address = $9
WHERE id = $1
RETURNING *;

-- name: DeleteAddress :exec
DELETE FROM address
WHERE id = $1;