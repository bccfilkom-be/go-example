// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package postgresql

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Book struct {
	ID        int64
	Title     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}