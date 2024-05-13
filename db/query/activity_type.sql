-- name: Getactivity_type :one
SELECT * FROM activity_type
WHERE id = $1 LIMIT 1;


-- name: Listactivity_types :many
SELECT * FROM activity_type
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: Createactivity_type :one
-- input: name 
-- output :one
INSERT INTO activity_type (
  name 
) VALUES (
  $1 
)
RETURNING *;

-- name: Updateactivity_type :one
UPDATE activity_type
  set name = $2
WHERE id = $1
RETURNING *;

-- name: Deleteactivity_type :exec
DELETE FROM activity_type
WHERE id = $1;


