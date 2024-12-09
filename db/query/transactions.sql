-- name: CreateTransaction :one
INSERT INTO transactions (
 tx_id,
 address_id,
 shipping_id,
 payment_method_id,
 shipping_price,
 total_price,
 status
) VALUES (
  $1, $2 ,$3, $4, $5, $6, $7
)
RETURNING *;


-- name: GetTransaction :one
SELECT 
t.tx_id as tx_id,
jsonb_build_object(
      'id', a.id,
      'recipient_name', a.recipient_name,
      'recipient_phone_number', a.recipient_phone_number,
      'province',a.province,
      'city',a.city,
      'district',a.district,
      'village',a.village,
      'postal_code',a.postal_code,
      'full_address',a.full_address      
) AS address,
jsonb_build_object(
      'id',s.id,
      'name',s.name,
      'price',s.price,
      'state',s.state,
      'updated_at',s.updated_at,
      'created_at',s.created_at
) AS shipping,
jsonb_build_object(
      'id',pm.id,
      'account_name',pm.account_name,
      'account_number',pm.account_number,
      'bank_name',pm.bank_name,
      'bank_images',pm.bank_images,
      'updated_at',pm.updated_at,
      'created_at',pm.created_at
) AS payment_method,
 FROM transactions t 
 LEFT JOIN address a ON t.address_id = a.id
 LEFT JOIN shippings s ON t.shipping_id = s.id
 LEFT JOIN payment_method pm ON t.payment_method_id = pm.id
WHERE tx_id = $1 LIMIT 1;

-- name: GetTransactionForUpdate :one
SELECT * FROM transactions
WHERE tx_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListTransaction :many
SELECT * FROM transactions
ORDER BY tx_id
LIMIT $1
OFFSET $2;

-- name: UpdateTransaction :one
UPDATE transactions
 set status = $2
WHERE tx_id = $1
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE tx_id = $1;