// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const countPets = `-- name: CountPets :one
SELECT
  COUNT(*)
FROM
  pets
`

func (q *Queries) CountPets(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countPets)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const listPets = `-- name: ListPets :many
SELECT
  id,
  category_id,
  name,
  photoURL,
  sold
FROM
  pets
LIMIT
  "$1::integer"
OFFSET
  "$2::integer"
`

type ListPetsRow struct {
	ID         int64
	CategoryID pgtype.Int8
	Name       string
	Photourl   string
	Sold       bool
}

func (q *Queries) ListPets(ctx context.Context) ([]ListPetsRow, error) {
	rows, err := q.db.Query(ctx, listPets)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPetsRow
	for rows.Next() {
		var i ListPetsRow
		if err := rows.Scan(
			&i.ID,
			&i.CategoryID,
			&i.Name,
			&i.Photourl,
			&i.Sold,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
