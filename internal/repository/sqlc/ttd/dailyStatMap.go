package ttd

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc"
)

func ToDomainDailyStat(row sqlc.DailyStat) domain.DailyStat {

	return domain.DailyStat{
		ID:         row.ID,
		CampaignID: row.CampaignID,
		DateDaily:  row.DateDaily,
		Impression: row.Impressions,
		Clicks:     row.Clicks,
		Cost:       row.Cost,
		Conversion: toLessPointerInt64(row.Conversion),
		Revenue:    toDecimal(row.Revenue),
		Reach:      row.Reach,
	}

}

func ToDomainDailyStatParams(limit int32, offset int32) sqlc.GetDailysParams {
	return sqlc.GetDailysParams{
		Limit:  limit,
		Offset: offset,
	}
}

func ToDomainDailyStats(rows []sqlc.DailyStat) []domain.DailyStat {
	out := make([]domain.DailyStat, len(rows))

	for i, row := range rows {
		out[i] = ToDomainDailyStat(row)
	}

	return out
}

func ToCreateDailyStatParams(ds domain.DailyStat) sqlc.CreateDailyStatParams {
	return sqlc.CreateDailyStatParams{
		CampaignID:  ds.CampaignID,
		DateDaily:   ds.DateDaily,
		Impressions: ds.Impression,
		Clicks:      ds.Clicks,
		Cost:        ds.Cost,
		Conversion:  toAddPointerInt64(ds.Conversion),
		Revenue:     toNullDecimalLessPoint(ds.Revenue),
		Reach:       ds.Reach,
	}
}

func ToUpdateDailyStat(id int64, dsp domain.DailyStatPath) sqlc.UpdateDailyParams {
	return sqlc.UpdateDailyParams{
		NewCampaignID: dsp.CampaignID,
		DateDaily:     dsp.DateDaily,
		Impressions:   dsp.Impression,
		Clicks:        dsp.Clicks,
		Cost:          toNullDecimal(dsp.Cost),
		Conversion:    dsp.Conversion,
		Revenue:       toNullDecimal(dsp.Revenue),
		Reach:         dsp.Reach,
		ID:            id,
	}
}
