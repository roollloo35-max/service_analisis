package service

// функция добовления компании в БД для создания сервиса
import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc/ttd"
	"context"
	"fmt"
)

func (sc *CampaignService) ServiceCreateCampaign(ctx context.Context, cid CID_Campaign) error {

	dc := domain.Campaign{
		Name:      cid.Name,
		Status:    domain.Status(cid.Status),
		Budget:    cid.Budget,
		TargetURL: cid.TargetURL,
		StartDate: cid.StartDate,
		EndDate:   cid.EndDate,
	}

	dcp := ttd.ToCreateParams(dc)

	_, err := sc.repoCreate.CreateCampaign(ctx, dcp)
	if err != nil {
		return fmt.Errorf("failed to create campaign: %w", err)
	}

	return nil

}
