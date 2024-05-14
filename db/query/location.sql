-- name: GetLocation :one
SELECT * FROM locations
WHERE id = $1 LIMIT 1;

-- name: GetLocationsByCountry :many
SELECT * FROM locations
WHERE country_id = $1;

-- name: ListLocations :many
SELECT * FROM locations
ORDER BY location_name
LIMIT $1
OFFSET $2;

-- name: CreateLocation :one
-- input: location_name, location_description, country_id
-- output :one
INSERT INTO locations (
  location_name, location_description, country_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateLocation :one
UPDATE locations
  set location_name = $2,
 location_description = $3,
 country_id = $4
WHERE id = $1
RETURNING *;

-- name: DeleteLocation :exec
DELETE FROM locations
WHERE id = $1;
