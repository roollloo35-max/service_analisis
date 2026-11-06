package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, os.Getenv("postgresql://postgres:fgtH123G@localhost:5432/DataForAnalisis"))
	if err != nil {
		panic(fmt.Sprint("unable to create connection pool: %v", err))
	}

	defer pool.Close()

	// queries := sqlc.New(pool)

}
