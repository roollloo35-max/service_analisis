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
		Amount:     toDecimal(row.Amount),
	}
}

func toCreateConversionEvent(dce domain.ConversionEvent) sqlc.ConversionEvent {
	return sqlc.ConversionEvent{
		Amount: toNullDecimalLessPoint(dce.Amount),
	}
}

func toUpdateConversionEvent(id int64, dcep domain.ConversionEventPath) sqlc.UpDataConversionParams {
	return sqlc.UpDataConversionParams{
		NewCampaignID: dcep.CampaignID,
		NewOccurredAt: dcep.OccurredAt,
		NewAmount:     toNullDecimal(dcep.Amount),
		ID:            id,
	}
}
