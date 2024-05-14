-- name: GetFood :one
SELECT * FROM food
WHERE id = $1 LIMIT 1;

-- name: GetfoodByLocation :many
SELECT * FROM food
WHERE area= $1;

-- name: Listfood :many
SELECT * FROM food
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: CreateFood :one
-- input: name, area, food_type, webaddress
-- output :one
INSERT INTO food (
  name, area, food_type, webaddress
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateFood :one
UPDATE food
  set name = $2,
 area   = $3,
 food_type = $4,
 webaddress = $5
WHERE id = $1
RETURNING *;

-- name: DeleteFood :exec
DELETE FROM food
WHERE id = $1;
