package ttd

import (
	"alisisService/internal/domain"

	"github.com/shopspring/decimal"
)

func StatusPtrToString(s *domain.Status) *string {
	if s == nil {
		return nil
	}
	v := string(*s)

	return &v
}

func toNullDecimal(d *decimal.Decimal) *decimal.NullDecimal {
	if d == nil {
		return nil
	}

	return &decimal.NullDecimal{
		Decimal: *d,
		Valid:   true,
	}
}

func toNullDecimalLessPoint(d decimal.Decimal) *decimal.NullDecimal {
	if d.Equal(decimal.Zero) {
		return nil
	}

	return &decimal.NullDecimal{
		Decimal: d,
		Valid:   true,
	}
}

func toDecimal(nd *decimal.NullDecimal) decimal.Decimal {
	if nd == nil || !nd.Valid {
		return decimal.Zero
	}

	return nd.Decimal
}

func toLessPointerInt64(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func toAddPointerInt64(i int64) *int64 {
	if i == 0 {
		return nil
	}
	return &i
}
