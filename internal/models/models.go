package models

import (
	"time"
)

type Campaining struct {
	ID           int
	Name         string
	Status       string
	CreatedAt    time.Time
	Budget       float64
	CostPerClick float64
	TargetURL    string
	StartDate    time.Time
	EndDate      time.Time
}

func NewCompaign(id int, name string, status string, budget float64, costPerClicl float64, targetURL string) *Campaining {
	return &Campaining{
		ID:           id,
		Name:         name,
		Status:       status,
		CreatedAt:    time.Now(),
		Budget:       budget,
		CostPerClick: costPerClicl,
	}
}

type Event struct {
	ID           int
	CampainingId int
	TypEvent     string
	Device       string
	IPAddress    string
	CreatedAt    time.Time
}

func NewEvent(id int, campaining_id int, typEvent string, device string, ip_adress string) *Event {
	return &Event{
		ID:           id,
		CampainingId: campaining_id,
		TypEvent:     typEvent,
		Device:       device,
		IPAddress:    ip_adress,
		CreatedAt:    time.Now(),
	}
}
