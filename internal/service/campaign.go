package service

import (
	"alisisService/internal/domain"
	"context"
)

type CampaignSecondary interface {
	Create(ctx context.Context, dc domain.Campaign) (domain.Campaign, error)
	GetList(ctx context.Context, offset int32, limit int32) ([]domain.Campaign, error)
	Get(ctx context.Context, id int64) (domain.Campaign, error)
	Update(ctx context.Context, id int64, dcp domain.CampaignPath) error
}

type RepoCamp struct {
	repo CampaignSecondary
}

func NewCampaignService(_repo CampaignSecondary) *RepoCamp {
	return &RepoCamp{
		repo: _repo,
	}
}
