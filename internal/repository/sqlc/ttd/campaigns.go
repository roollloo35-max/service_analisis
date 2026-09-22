package ttd

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc"
	"time"
)

func toDomaintCampaign(row sqlc.Campaign) domain.Campaign {
	var endDate *time.Time

	if row.EndDate.Valid {
		t := row.EndDate.Time
		endDate = &t
	}

	return domain.Campaign{
		ID:         row.ID,
		Status:     domain.Status(row.Status),
		Budget:     row.Budget,
		CreatedAt:  row.CreatedAt,
		TargetURL:  row.TargetUrl,
		StartDate:  row.StartDate,
		EndDate:    endDate,
		LastUpdate: row.LastUpdate.Time,
	}
}

func toCreateParams(c domain.Campaign) sqlc.CreateCampaignParams {
	return sqlc.CreateCampaignParams{}
}

func toUpdateParams(id int64, p domain.CampaignPatch) sqlc.UpdateCampaignParams {
	return sqlc.UpdateCampaignParams{}
}
