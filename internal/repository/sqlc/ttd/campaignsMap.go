package ttd

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc"
)

func ToCampaignRow(row sqlc.CreateCampaignRow) sqlc.Campaign {
	return sqlc.Campaign{
		ID:         row.ID,
		Name:       row.Name,
		Status:     row.Status,
		Budget:     row.Budget,
		TargetUrl:  row.TargetUrl,
		StartDate:  row.StartDate,
		EndDate:    row.EndDate,
		LastUpdate: row.EndDate,
		ExternalID: row.ExternalID,
		DeleteAtr:  row.DeleteAtr,
		VersionAtr: row.VersionAtr,
	}
}

func ToGetCampaign(row sqlc.GetCampaignRow) domain.Campaign {

	return domain.Campaign{
		ID:         row.ID,
		Status:     domain.Status(row.Status),
		Budget:     row.Budget,
		CreatedAt:  row.CreatedAt,
		TargetURL:  row.TargetUrl,
		StartDate:  row.StartDate,
		EndDate:    row.EndDate,
		LastUpdate: row.LastUpdate,
		ExternalID: row.ExternalID,
		DeleteAtr:  row.DeleteAtr,
		VersionAtr: row.VersionAtr,
	}
}

func ToListCampaignParams(offset int32, limit int32) sqlc.GetAllCampaignsParams {
	return sqlc.GetAllCampaignsParams{
		Offset: offset,
		Limit:  limit,
	}
}

func ToListCampaign(rows []sqlc.GetAllCampaignsRow) []domain.Campaign {

	rowsList := make([]domain.Campaign, 0, len(rows))

	for _, row := range rows {
		rowsList = append(rowsList, domain.Campaign{
			ID:         row.ID,
			Name:       row.Name,
			Status:     domain.Status(row.Status),
			Budget:     row.Budget,
			CreatedAt:  row.CreatedAt,
			TargetURL:  row.TargetUrl,
			StartDate:  row.StartDate,
			EndDate:    row.EndDate,
			LastUpdate: row.LastUpdate,
			ExternalID: row.ExternalID,
			DeleteAtr:  row.DeleteAtr,
			VersionAtr: row.VersionAtr,
		})
	}
	return rowsList
}

// перепроверить алгоритм

func ToCreateParams(c domain.Campaign) sqlc.CreateCampaignParams {
	return sqlc.CreateCampaignParams{
		Name:       c.Name,
		Status:     string(c.Status),
		Budget:     c.Budget,
		TargetUrl:  c.TargetURL,
		StartDate:  c.StartDate,
		EndDate:    c.EndDate,
		ExternalID: c.ExternalID,
		DeleteAtr:  c.DeleteAtr,
		VersionAtr: c.VersionAtr,
	}

}

func ToUpdateParams(id int64, p domain.CampaignPath) sqlc.UpdateCampaignParams {

	return sqlc.UpdateCampaignParams{
		NewName:       p.Name,
		NewStatus:     StatusPtrToString(p.Status),
		NewBudget:     toNullDecimal(p.Budget),
		NewTargetUrl:  p.TargetURL,
		NewStartDate:  p.StartDate,
		NewEndDate:    p.EndDate,
		ID:            id,
		NewExternalID: p.ExternalID,
		NewDeleteAtr:  p.DeleteAtr,
		NewVersionAtr: p.VersionAtr,
	}
}
