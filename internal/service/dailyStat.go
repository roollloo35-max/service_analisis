package service

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc"
	"alisisService/internal/repository/sqlc/ttd"
	"context"
	"fmt"
)

type SecondDailyStat interface {
	CreateDailyStat(ctx context.Context, arg sqlc.CreateDailyStatParams) (sqlc.DailyStat, error)
	GetDaily(ctx context.Context, id int64) (sqlc.DailyStat, error)
	GetDailys(ctx context.Context, arg sqlc.GetDailysParams) ([]sqlc.DailyStat, error)
	UpdateDaily(ctx context.Context, arg sqlc.UpdateDailyParams) (sqlc.DailyStat, error)
}

type RepoDaily struct {
	repo SecondDailyStat
}

func NewCampaignSevice(_repo SecondDailyStat) *RepoDaily {
	return &RepoDaily{
		repo: _repo,
	}
}

func (sd *RepoDaily) GetListDaily(ctx context.Context, cid CIDPaginDaily) ([]domain.DailyStat, error) {
	params := ttd.ToDomainDailyStatParams(cid.Limit, cid.Offset)

	listDailyRow, err := sd.repo.GetDailys(ctx, params)

	if err != nil {
		return nil, fmt.Errorf("get list daily stat %w", err)
	}

	listDaily := ttd.ToDomainDailyStats(listDailyRow)

	return listDaily, nil
}

func (sd *RepoDaily) GetDaily(ctx context.Context, cid CIDDailyStat) (domain.DailyStat, error) {

	listDailyRow, err := sd.repo.GetDaily(ctx, cid.ID)

	if err != nil {
		return domain.DailyStat{}, fmt.Errorf("get list daily stat %w", err)
	}

	Daily := ttd.ToDomainDailyStat(listDailyRow)

	return Daily, nil
}
