package service

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc/ttd"
	"context"
	"fmt"
)

// func (sc *CampaignService) ServiceGetAllCampaign(ctx context.Context) ([]domain.Campaign, error) {
// 	rows := make([]domain.Campaign, 0)

// 	rows, err :=

// 	return nil
// }

func (sc *CampaignService) ServiceGetCampaign(ctx context.Context, id int64) (domain.Campaign, error) {

	var dc domain.Campaign

	oldCampaignRow, err := sc.repoGet.GetCampaign(ctx, id)

	if err != nil {
		return dc, fmt.Errorf("Not found campaign: %w", err)
	}

	dc = ttd.ToGetCampaign(oldCampaignRow)

	return dc, nil
}
