-- name: GetAccommodation_type :one
SELECT * FROM accommodation_type
WHERE id = $1 LIMIT 1;


-- name: ListAccommodation_types :many
SELECT * FROM accommodation_type
ORDER BY type
LIMIT $1
OFFSET $2;

-- name: CreateAccommodation_type :one
-- input: type 
-- output :one
INSERT INTO accommodation_type (
  type 
) VALUES (
  $1 
)
RETURNING *;

-- name: UpdateAccommodation_type :one
UPDATE accommodation_type
  set type = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccommodation_type :exec
DELETE FROM accommodation_type
WHERE id = $1;


