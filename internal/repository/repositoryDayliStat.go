package repository

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc"
	"alisisService/internal/repository/sqlc/ttd"
	"context"
)

type DayliStatRepository interface {
	UpdateDaily(ctx context.Context, arg sqlc.UpdateDailyParams) (sqlc.DailyStat, error)
	GetDailys(ctx context.Context, arg sqlc.GetDailysParams) ([]sqlc.DailyStat, error)
	GetDaily(ctx context.Context, id int64) (sqlc.DailyStat, error)
	CreateDailyStat(ctx context.Context, arg sqlc.CreateDailyStatParams) (sqlc.DailyStat, error)
}

type RepoDailyStat struct {
	Repo DayliStatRepository
}

func NewDayliStatRepo(repo DayliStatRepository) *RepoDailyStat {
	return &RepoDailyStat{
		Repo: repo,
	}
}

func (rds *RepoDailyStat) Create(ctx context.Context, dcp domain.DailyStat) (domain.DailyStat, error) {

	dsr := ttd.ToCreateDailyStatParams(dcp)

	dailyStatRow, err := rds.Repo.CreateDailyStat(ctx, dsr)

	if err != nil {
		return domain.DailyStat{}, err
	}

	return ttd.ToDomainDailyStat(dailyStatRow), nil
}

func (rds *RepoDailyStat) Get(ctx context.Context, id int64) (domain.DailyStat, error) {

	dailyStatRow, err := rds.Repo.GetDaily(ctx, id)
	if err != nil {
		return domain.DailyStat{}, err
	}

	return ttd.ToDomainDailyStat(dailyStatRow), err

}

func (rds *RepoDailyStat) GetList(ctx context.Context, offset int32, limit int32) ([]domain.DailyStat, error) {

	dailysParams := ttd.ToDomainDailyStatParams(limit, offset)

	LsitDailyStat, err := rds.Repo.GetDailys(ctx, dailysParams)

	if err != nil {
		return nil, err
	}

	return ttd.ToDomainDailyStats(LsitDailyStat), nil

}

func (rds *RepoDailyStat) Update(ctx context.Context, id int64, dcp domain.DailyStatPath) (domain.DailyStat, error) {

	dcpu := ttd.ToUpdateDailyStat(id, dcp)

	dailyStatRow, err := rds.Repo.UpdateDaily(ctx, dcpu)

	if err != nil {
		return domain.DailyStat{}, err
	}

	return ttd.ToDomainDailyStat(dailyStatRow), nil

}
