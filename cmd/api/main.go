package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbPool, err := pgxpool.New(context.Background(), "postgres://username:password@localhost:5432/mydb")
	if err != nil {
		log.Fatal("Failed to make db connection %v", err)
	}
	defer dbPool.Close()

}
