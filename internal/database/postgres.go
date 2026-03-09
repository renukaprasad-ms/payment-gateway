package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(url string) *pgxpool.Pool {

	db, err := pgxpool.New(context.Background(), url)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
