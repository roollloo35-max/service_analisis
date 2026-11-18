package repository

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc"
	"alisisService/internal/repository/sqlc/ttd"

	"context"
)

type ConversionEventRepository interface {
	CreateConversionEvent(ctx context.Context, arg sqlc.CreateConversionEventParams) (sqlc.ConversionEvent, error)
	GetConversion(ctx context.Context, id int64) (sqlc.ConversionEvent, error)
	UpDataConversion(ctx context.Context, arg sqlc.UpDataConversionParams) (sqlc.ConversionEvent, error)
}

type RepoConversionEvent struct {
	Repo ConversionEventRepository
}

func (ce *RepoConversionEvent) Create(ctx context.Context, d domain.ConversionEvent) (domain.ConversionEvent, error) {
	dcp := ttd.ToCreateConversionEvent(d)

	dcpp := ttd.ToConversionParams(dcp)

	conversionEvent, err := ce.Repo.CreateConversionEvent(ctx, dcpp)
	if err != nil {
		return domain.ConversionEvent{}, err
	}

	convertToDomain := ttd.ToDomainConversionEvent(conversionEvent)

	return convertToDomain, nil

}

func (ce *RepoConversionEvent) Update(ctx context.Context, id int64, d domain.ConversionEventPath) (domain.ConversionEvent, error) {
	dcp := ttd.ToUpdateConversionEvent(id, d)

	conversionEventRow, err := ce.Repo.UpDataConversion(ctx, dcp)
	if err != nil {
		return domain.ConversionEvent{}, err
	}

	convertToDomain := ttd.ToDomainConversionEvent(conversionEventRow)

	return convertToDomain, nil
}
