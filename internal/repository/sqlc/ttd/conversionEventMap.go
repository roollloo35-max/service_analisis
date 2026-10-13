package ttd

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc"
)

func toDomainConversionEvent(row sqlc.ConversionEvent) domain.ConversionEvent {
	return domain.ConversionEvent{
		ID:         row.ID,
		CampaignID: row.CampaignID,
		OccurredAt: row.OccurredAt,
		Amount:     row.Amount,
	}
}

func toCreateConversionEvent(dce domain.ConversionEvent) sqlc.ConversionEvent {
	return sqlc.ConversionEvent{
		Amount: dce.Amount,
	}
}

func toUpdateConversionEvent(id int64, dcep domain.ConversionEventPath) sqlc.UpDataConversionParams {
	return sqlc.UpDataConversionParams{
		NewCampaignID: toNullInt64(dcep.CampaignID),
		NewOccurredAt: toNullTime(dcep.OccurredAt),
		NewAmount:     toNullDecimal(dcep.Amount),
		ID:            id,
	}
}
