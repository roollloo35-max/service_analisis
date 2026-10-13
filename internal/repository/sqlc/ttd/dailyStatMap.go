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
		Conversion: row.Conversion.Int32,
		Revenue:    row.Revenue,
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
		Conversion:  toNullInt32(ds.Conversion),
		Revenue:     ds.Revenue,
		Reach:       ds.Reach,
	}
}

func ToUpdateDailyStat(id int64, dsp domain.DailyStatPath) sqlc.UpdateDailyParams {
	return sqlc.UpdateDailyParams{
		NewCampaignID: toNullInt64(dsp.CampaignID),
		DateDaily:     toNullTime(dsp.DateDaily),
		Impressions:   toNullInt64(dsp.Impression),
		Clicks:        toNullInt64(dsp.Clicks),
		Cost:          toNullDecimal(dsp.Cost),
		Conversion:    toNullInt32IsPointner(dsp.Conversion),
		Revenue:       toNullDecimal(dsp.Revenue),
		Reach:         toNullInt64(dsp.Reach),
		ID:            id,
	}
}
