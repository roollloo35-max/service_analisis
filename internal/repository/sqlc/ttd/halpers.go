package ttd

import (
	"alisisService/internal/domain"
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
)

func toNullDecimal(x *decimal.Decimal) decimal.NullDecimal {
	if x == nil {
		return decimal.NullDecimal{Decimal: *x, Valid: false}
	}

	return decimal.NullDecimal{
		Decimal: *x,
		Valid:   true,
	}
}

func toNullString(x *string) sql.NullString {
	if x == nil {
		return sql.NullString{Valid: false}
	}

	return sql.NullString{
		String: *x,
		Valid:  true,
	}
}

func toNullStatus(x *domain.Status) sql.NullString {

	if x == nil {
		return sql.NullString{Valid: false}
	}

	return sql.NullString{
		String: string(*x),
		Valid:  true,
	}
}

func toNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{
		Time:  *t,
		Valid: true}
}

func toNullInt32(cons int32) sql.NullInt32 {
	if cons == 0 {
		return sql.NullInt32{Int32: cons, Valid: false}
	}

	return sql.NullInt32{
		Int32: cons,
		Valid: true,
	}
}
func toNullInt32IsPointner(cons *int32) sql.NullInt32 {
	if cons == nil {
		return sql.NullInt32{Int32: *cons, Valid: false}
	}

	return sql.NullInt32{
		Int32: *cons,
		Valid: true,
	}
}

func toNullInt64(cons *int64) sql.NullInt64 {
	if cons == nil {
		return sql.NullInt64{Int64: *cons, Valid: false}
	}

	return sql.NullInt64{
		Int64: *cons,
		Valid: true,
	}
}
