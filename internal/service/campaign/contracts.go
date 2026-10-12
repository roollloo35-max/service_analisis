package service

import (
	"time"

	"github.com/shopspring/decimal"
)

type CID_Campaign struct {
	Name       string
	Status     string
	Budget     decimal.Decimal
	CreatedAt  time.Time
	TargetURL  string
	StartDate  time.Time
	EndDate    *time.Time
	LastUpdate time.Time
}

// adapter with handler
