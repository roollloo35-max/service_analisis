package ttd

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc"
)

func ToDomainConversionEvent(row sqlc.ConversionEvent) domain.ConversionEvent {
	return domain.ConversionEvent{
		ID:         row.ID,
		CampaignID: row.CampaignID,
		OccurredAt: row.OccurredAt,
		Amount:     toDecimal(row.Amount),
	}
}

func ToConversionParams(row sqlc.ConversionEvent) sqlc.CreateConversionEventParams {
	return sqlc.CreateConversionEventParams{
		CampaignID: row.CampaignID,
		Amount:     row.Amount,
	}
}

func ToCreateConversionEvent(dce domain.ConversionEvent) sqlc.ConversionEvent {
	return sqlc.ConversionEvent{
		Amount:     toNullDecimalLessPoint(dce.Amount),
		CampaignID: dce.CampaignID,
	}
}

func ToUpdateConversionEvent(id int64, dcep domain.ConversionEventPath) sqlc.UpDataConversionParams {
	return sqlc.UpDataConversionParams{
		NewCampaignID: dcep.CampaignID,
		NewOccurredAt: dcep.OccurredAt,
		NewAmount:     toNullDecimal(dcep.Amount),
		ID:            id,
	}
}
