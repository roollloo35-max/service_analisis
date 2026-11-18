package repository

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc"
	"alisisService/internal/repository/sqlc/ttd"
	"context"
	"fmt"
)

type AdPlatformRepository interface {
	CreatePlatform(ctx context.Context, arg sqlc.CreatePlatformParams) (sqlc.AdPlatform, error)
	GetAdPlatform(ctx context.Context, id int64) (sqlc.AdPlatform, error)
	GetListAdPlatform(ctx context.Context, arg sqlc.GetListAdPlatformParams) ([]sqlc.AdPlatform, error)
	UpdateAdPlatform(ctx context.Context, arg sqlc.UpdateAdPlatformParams) (sqlc.AdPlatform, error)
}

type AdPlatformRepo struct {
	Repo AdPlatformRepository
}

func (ad *AdPlatformRepo) Create(ctx context.Context, d domain.AdPlatform) (domain.AdPlatform, error) {

	dsr := ttd.ToCreateAdParams(d)

	platformRow, err := ad.Repo.CreatePlatform(ctx, dsr)
	if err != nil {
		return domain.AdPlatform{}, fmt.Errorf("erros create row", err)
	}

	dom := ttd.ToDomain(platformRow)

	return dom, nil

}

func (ad *AdPlatformRepo) Get(ctx context.Context, id int64) (domain.AdPlatform, error) {

	oldPlatform, err := ad.Repo.GetAdPlatform(ctx, id)
	if err != nil {
		return domain.AdPlatform{}, fmt.Errorf("errors get", err)
	}

	dp := ttd.ToDomain(oldPlatform)

	return dp, nil
}
