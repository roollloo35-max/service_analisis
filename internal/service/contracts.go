package service

import (
	"time"

	"github.com/shopspring/decimal"
)

type CIDCampaign struct {
	Name       string
	Status     string
	Budget     decimal.Decimal
	CreatedAt  time.Time
	TargetURL  string
	StartDate  time.Time
	EndDate    *time.Time
	LastUpdate time.Time
}

type CIDPaginCamp struct {
	Offset int32
	Limit  int32
}

// adapter with handler

type CIDDailyStat struct {
	ID          int64
	CampagignID int64
	DateDaily   time.Time
	Impression  int64
	Clicks      int64
	Cost        decimal.Decimal
	Conversion  int32
	Revenue     decimal.Decimal
	Reach       int64
}

type CIDPaginDaily struct {
	Offset int32
	Limit  int32
}
