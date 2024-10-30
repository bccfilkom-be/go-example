package app

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewPostgreSQL(url string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	return conn
}
