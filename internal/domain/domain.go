package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type Status string

const (
	StatusDraft    Status = "draft"
	StatusActive   Status = "active"
	StatusPaused   Status = "pause"
	StatusArchived Status = "archived"
)

type Campaign struct {
	ID         int64
	Name       string
	Status     Status // enum
	Budget     decimal.Decimal
	CreatedAt  time.Time
	TargetURL  string
	StartDate  time.Time
	EndDate    *time.Time
	LastUpdate time.Time
}

type CampaignPath struct {
	Name      *string
	Status    *Status // enaum
	Budget    *decimal.Decimal
	TargetURL *string
	StartDate *time.Time
	EndDate   *time.Time
}

type DailyStat struct {
	ID         int64
	CampaignID int64
	DateDaily  time.Time
	Impression int64
	Clicks     int64
	Cost       decimal.Decimal
	Conversion int32
	Revenue    decimal.Decimal
	Reach      int64
}

type DailyStatPath struct {
	CampaignID *int64
	DateDaily  *time.Time
	Impression *int64
	Clicks     *int64
	Cost       *decimal.Decimal
	Conversion *int32
	Revenue    *decimal.Decimal
	Reach      *int64
}

type ConversionEvent struct {
	ID         int64
	CampaignID int64
	OccurredAt time.Time
	Amount     decimal.Decimal
}

type ConversionEventPath struct {
	CampaignID *int64
	OccurredAt *time.Time
	Amount     *decimal.Decimal
}
