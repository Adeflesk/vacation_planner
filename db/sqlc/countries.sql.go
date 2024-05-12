// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: countries.sql

package db

import (
	"context"
)

const createCountry = `-- name: CreateCountry :one
INSERT INTO countries (
  name, continent_name
) VALUES (
  $1, $2
)
RETURNING id, name, continent_name
`

type CreateCountryParams struct {
	Name          string `db:"name"`
	ContinentName string `db:"continent_name"`
}

// input: name, continent_name
// output :one
func (q *Queries) CreateCountry(ctx context.Context, arg CreateCountryParams) (Country, error) {
	row := q.db.QueryRowContext(ctx, createCountry, arg.Name, arg.ContinentName)
	var i Country
	err := row.Scan(&i.ID, &i.Name, &i.ContinentName)
	return i, err
}

const deleteCountry = `-- name: DeleteCountry :exec
DELETE FROM countries
WHERE id = $1
`

func (q *Queries) DeleteCountry(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCountry, id)
	return err
}

const getCountry = `-- name: GetCountry :one
SELECT id, name, continent_name FROM countries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCountry(ctx context.Context, id int64) (Country, error) {
	row := q.db.QueryRowContext(ctx, getCountry, id)
	var i Country
	err := row.Scan(&i.ID, &i.Name, &i.ContinentName)
	return i, err
}

const listCountries = `-- name: ListCountries :many
SELECT id, name, continent_name FROM countries
ORDER BY name
`

func (q *Queries) ListCountries(ctx context.Context) ([]Country, error) {
	rows, err := q.db.QueryContext(ctx, listCountries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Country{}
	for rows.Next() {
		var i Country
		if err := rows.Scan(&i.ID, &i.Name, &i.ContinentName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCountry = `-- name: UpdateCountry :exec
UPDATE countries
  set name = $2,
  continent_name = $3
WHERE id = $1
`

type UpdateCountryParams struct {
	ID            int64  `db:"id"`
	Name          string `db:"name"`
	ContinentName string `db:"continent_name"`
}

func (q *Queries) UpdateCountry(ctx context.Context, arg UpdateCountryParams) error {
	_, err := q.db.ExecContext(ctx, updateCountry, arg.ID, arg.Name, arg.ContinentName)
	return err
}