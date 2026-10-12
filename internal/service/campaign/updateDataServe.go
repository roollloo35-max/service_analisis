package service

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc/ttd"
	"context"
	"fmt"
)

func (sc *CampaignService) ServiceUpdateCampaign(ctx context.Context, id int64, cid CID_Campaign) error {

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

	_, err := sc.repoUpdate.UpdateCampaign(ctx, dcp)

	if err != nil {
		return fmt.Errorf("failed to update campagn %d: %w", id, err)
	}

	return nil
}
