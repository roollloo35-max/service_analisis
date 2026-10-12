package service

import (
	"alisisService/internal/repository/sqlc"
	"context"
)

type SecondaryCreateCampaign interface {
	CreateCampaign(ctx context.Context, arg sqlc.CreateCampaignParams) (sqlc.CreateCampaignRow, error)
}

type SecondaryUpdateCampaign interface {
	UpdateCampaign(ctx context.Context, arg sqlc.UpdateCampaignParams) (sqlc.UpdateCampaignRow, error)
}

type SecondaryGetCampaign interface {
	GetCampaign(ctx context.Context, id int64) (sqlc.GetCampaignRow, error)
}

type SecondaryGetAllCampaign interface {
	GetAllCampaign(ctx context.Context) ([]sqlc.GetAllCampaignsRow, error)
}

type CampaignService struct {
	repoCreate SecondaryCreateCampaign
	repoUpdate SecondaryUpdateCampaign
	repoGet    SecondaryGetCampaign
	repoGetAll SecondaryGetAllCampaign
}

func NewCampaignService(_repoCreate SecondaryCreateCampaign, _repoUpdate SecondaryUpdateCampaign, _repoGet SecondaryGetCampaign, _repoGetAll SecondaryGetAllCampaign) *CampaignService {
	return &CampaignService{
		repoCreate: _repoCreate,
		repoUpdate: _repoUpdate,
		repoGet:    _repoGet,
		repoGetAll: _repoGetAll,
	}
}
