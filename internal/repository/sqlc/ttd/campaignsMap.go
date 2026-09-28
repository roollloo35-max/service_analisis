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

// перепроверить алгоритм

func toCreateParams(c domain.Campaign) sqlc.CreateCampaignParams {
	return sqlc.CreateCampaignParams{
		Name:      c.Name,
		Status:    string(c.Status),
		Budget:    c.Budget,
		TargetUrl: c.TargetURL,
		StartDate: c.StartDate,
		EndDate:   toNullTime(c.EndDate),
	}

}

func toUpdateParams(id int64, p domain.CampaignPath) sqlc.UpdateCampaignParams {

	return sqlc.UpdateCampaignParams{
		NewName:      toNullString(p.Name),
		NewStatus:    toNullStatus(p.Status),
		NewBudget:    toNullDecimal(p.Budget),
		NewTargetUrl: toNullString(p.TargetURL),
		NewStartDate: toNullTime(p.StartDate),
		NewEndDate:   toNullTime(p.EndDate),
		ID:           id,
	}
}
