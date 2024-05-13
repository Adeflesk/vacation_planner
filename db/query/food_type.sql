-- name: Getfood_type :one
SELECT * FROM food_type
WHERE id = $1 LIMIT 1;


-- name: Listfood_types :many
SELECT * FROM food_type
ORDER BY type
LIMIT $1
OFFSET $2;

-- name: Createfood_type :one
-- input: type 
-- output :one
INSERT INTO food_type (
  type 
) VALUES (
  $1 
)
RETURNING *;

-- name: Updatefood_type :one
UPDATE food_type
  set type = $2
WHERE id = $1
RETURNING *;

-- name: Deletefood_type :exec
DELETE FROM food_type
WHERE id = $1;

