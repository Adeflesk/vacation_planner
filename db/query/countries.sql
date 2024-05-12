-- name: GetCountry :one
SELECT * FROM countries
WHERE id = $1 LIMIT 1;

-- name: ListCountries :many
SELECT * FROM countries
ORDER BY name;

-- name: CreateCountry :one
-- input: name, continent_name
-- output :one
INSERT INTO countries (
  name, continent_name
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateCountry :exec
UPDATE countries
  set name = $2,
  continent_name = $3
WHERE id = $1;

-- name: DeleteCountry :exec
DELETE FROM countries
WHERE id = $1;