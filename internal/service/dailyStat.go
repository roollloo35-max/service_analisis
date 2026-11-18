package service

import (
	"alisisService/internal/domain"
	"context"
)

type SecondDailyStat interface {
	Create(ctx context.Context, dcp domain.DailyStat) (domain.DailyStat, error)
	Get(ctx context.Context, id int64) (domain.DailyStat, error)
	GetList(ctx context.Context, offset int32, limit int32) ([]domain.DailyStat, error)
	Update(ctx context.Context, id int64, dcp domain.DailyStatPath) (domain.DailyStat, error)
}

type RepoDaily struct {
	repo SecondDailyStat
}

func NewCampaignSevice(_repo SecondDailyStat) *RepoDaily {
	return &RepoDaily{
		repo: _repo,
	}
}
