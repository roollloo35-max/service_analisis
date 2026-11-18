package main

import (
	"fmt"
	"time"
)

func main() {

	// ctx := context.Background()

	// pool, err := pgxpool.New(ctx, os.Getenv("postgresql://postgres:fgtH123G@localhost:5432/DataForAnalisis"))
	// if err != nil {
	// 	panic(fmt.Sprint("unable to create connection pool: %v", err))
	// }

	// defer pool.Close()

	// // queries := sqlc.New(pool)

	fmt.Println("Hello, World!")

	time.Sleep(time.Second * 500000)

}
