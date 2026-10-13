package service

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc"
	"alisisService/internal/repository/sqlc/ttd"
	"context"
	"fmt"
)

type CampaignSecondary interface {
	CreateCampaign(ctx context.Context, arg sqlc.CreateCampaignParams) (sqlc.CreateCampaignRow, error)
	UpdateCampaign(ctx context.Context, arg sqlc.UpdateCampaignParams) (sqlc.UpdateCampaignRow, error)
	GetCampaign(ctx context.Context, id int64) (sqlc.GetCampaignRow, error)
	GetAllCampaign(ctx context.Context, arg sqlc.GetAllCampaignsParams) ([]sqlc.GetAllCampaignsRow, error)
}

type RepoCamp struct {
	repo CampaignSecondary
}

func NewCampaignService(_repo CampaignSecondary) *RepoCamp {
	return &RepoCamp{
		repo: _repo,
	}
}

func (sc *RepoCamp) GetListCampaigns(ctx context.Context, cid CIDPaginCamp) ([]domain.Campaign, error) {
	paramsRow := ttd.ToListCampaignParams(cid.Offset, cid.Limit)

	listCampRow, err := sc.repo.GetAllCampaign(ctx, paramsRow)

	if err != nil {
		return nil, fmt.Errorf("Get list campaigns: %w", err)
	}

	listCamp := ttd.ToListCampaign(listCampRow)

	return listCamp, nil
}

func (sc *RepoCamp) GetCampaign(ctx context.Context, id int64) (domain.Campaign, error) {

	var dc domain.Campaign

	oldCampaignRow, err := sc.repo.GetCampaign(ctx, id)

	if err != nil {
		return domain.Campaign{}, fmt.Errorf("Get campaign %d: %w", id, err)
	}

	dc = ttd.ToGetCampaign(oldCampaignRow)

	return dc, nil
}

func (sc *RepoCamp) CreateCampaign(ctx context.Context, cid CIDCampaign) (domain.Campaign, error) {

	dc := domain.Campaign{
		Name:      cid.Name,
		Status:    domain.Status(cid.Status),
		Budget:    cid.Budget,
		TargetURL: cid.TargetURL,
		StartDate: cid.StartDate,
		EndDate:   cid.EndDate,
	}

	dcp := ttd.ToCreateParams(dc)

	newRow, err := sc.repo.CreateCampaign(ctx, dcp)
	if err != nil {
		return domain.Campaign{}, fmt.Errorf("create comapign: %w", err)
	}

	camp := ttd.ToCampaignRow(newRow)

	domRow := ttd.ToDomaintCampaign(camp)

	return domRow, nil

}

func (sc *RepoCamp) UpdateCampaign(ctx context.Context, id int64, cid CIDCampaign) error {

	st := domain.Status(cid.Status)

	dc := domain.CampaignPath{
		Name:      &cid.Name,
		Status:    &st,
		Budget:    &cid.Budget,
		TargetURL: &cid.TargetURL,
		StartDate: &cid.StartDate,
		EndDate:   cid.EndDate,
	}

	dcp := ttd.ToUpdateParams(id, dc)

	_, err := sc.repo.UpdateCampaign(ctx, dcp)

	if err != nil {
		return fmt.Errorf("update campaign %d: %w", id, err)
	}

	return nil
}
