package repository

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc"
	"alisisService/internal/repository/sqlc/ttd"
	"context"
	"fmt"
)

type CampaginRepository interface {
	CreateCampaign(ctx context.Context, arg sqlc.CreateCampaignParams) (sqlc.CreateCampaignRow, error)
	UpdateCampaign(ctx context.Context, arg sqlc.UpdateCampaignParams) (sqlc.UpdateCampaignRow, error)
	GetCampaign(ctx context.Context, id int64) (sqlc.GetCampaignRow, error)
	GetAllCampaign(ctx context.Context, arg sqlc.GetAllCampaignsParams) ([]sqlc.GetAllCampaignsRow, error)
}

type RepoCampaign struct {
	Repo CampaginRepository
}

func NewCampaignRepo(repo CampaginRepository) *RepoCampaign {
	return &RepoCampaign{
		Repo: repo,
	}
}

func (rc *RepoCampaign) Create(ctx context.Context, dc domain.Campaign) (domain.Campaign, error) {

	dcp := ttd.ToCreateParams(dc)

	CampaignRow, err := rc.Repo.CreateCampaign(ctx, dcp)
	if err != nil {
		return domain.Campaign{}, fmt.Errorf("create campaign: %w", err)
	}

	camp := ttd.ToCampaignRow(CampaignRow)

	domainRow := ttd.ToDomaintCampaign(camp)

	return domainRow, nil

}

func (sc *RepoCampaign) GetList(ctx context.Context, offset int32, limit int32) ([]domain.Campaign, error) {
	paramsRow := ttd.ToListCampaignParams(offset, limit)

	listCampRow, err := sc.Repo.GetAllCampaign(ctx, paramsRow)

	if err != nil {
		return nil, fmt.Errorf("Get list campaigns: %w", err)
	}

	listCamp := ttd.ToListCampaign(listCampRow)

	return listCamp, nil
}

func (sc *RepoCampaign) Get(ctx context.Context, id int64) (domain.Campaign, error) {

	var dc domain.Campaign

	oldCampaignRow, err := sc.Repo.GetCampaign(ctx, id)

	if err != nil {
		return domain.Campaign{}, fmt.Errorf("Get campaign %d: %w", id, err)
	}

	dc = ttd.ToGetCampaign(oldCampaignRow)

	return dc, nil
}

func (sc *RepoCampaign) Update(ctx context.Context, id int64, dcp domain.CampaignPath) error {

	dcpu := ttd.ToUpdateParams(id, dcp)

	_, err := sc.Repo.UpdateCampaign(ctx, dcpu)

	if err != nil {
		return fmt.Errorf("update campaign %d: %w", id, err)
	}

	return nil
}
