-- name: GetAccommodation :one
SELECT * FROM Accommodation
WHERE id = $1 LIMIT 1;

-- name: GetAccommodationByLocation :many
SELECT * FROM Accommodation
WHERE area= $1;

-- name: ListAccommodation :many
SELECT * FROM Accommodation
ORDER BY accommodation_name
LIMIT $1
OFFSET $2;

-- name: CreateAccommodation :one
-- input: name, pernight, accommodation_type, accommodation_description, webaddress, emailadddress, phonenumber, area
-- output :one
INSERT INTO Accommodation (
  accommodation_name, pernight, accommodation_type, accommodation_description, webaddress,  emailaddress, phonenumber, area
) VALUES (
  $1, $2, $3, $4,$5, $6, $7, $8
)
RETURNING *;

-- name: UpdateAccommodation :one
UPDATE Accommodation
  set  accommodation_name = $2,
 pernight   = $3,
 Accommodation_type = $4,
 webaddress = $5,
 emailaddress = $6,
 phonenumber = $7,
 area = $8

WHERE id = $1
RETURNING *;

-- name: DeleteAccommodation :exec
DELETE FROM Accommodation
WHERE id = $1;
