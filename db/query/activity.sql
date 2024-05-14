-- name: GetActivity :one
SELECT * FROM Activity
WHERE id = $1 LIMIT 1;

-- name: GetActivityByLocation :many
SELECT * FROM Activity
WHERE area= $1;

-- name: ListActivity :many
SELECT * FROM Activity
ORDER BY activity_name 
LIMIT $1
OFFSET $2;

-- name: CreateActivity :one
-- input: activity_name, activity_type, description, webaddress, time_allocated, area
-- output :one
INSERT INTO Activity (
activity_name,activity_type,description,webaddress,time_allocated,area
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateActivity :one
UPDATE Activity
  set activity_name = $2,
  activity_type = $3,
  description = $4,
  webaddress = $5,
  time_allocated = $6,
  area = $7
WHERE id = $1
RETURNING *;

-- name: DeleteActivity :exec
DELETE FROM Activity
WHERE id = $1;
