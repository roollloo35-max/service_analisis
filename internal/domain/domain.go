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

type CampaignPatch struct {
	Name      *string
	Status    *Status // enaum
	Budget    *decimal.Decimal
	TargetURL *string
	StartDate *time.Time
	EndDate   *time.Time
}
