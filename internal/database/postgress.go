package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connecton(ctx context.Context) (*pgx.Conn, error) {

	return pgx.Connect(ctx, "postgresql://postgres:fgtH123G@localhost:5432/DataForAnalisis")

}
