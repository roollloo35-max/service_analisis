package ttd

import (
	"alisisService/internal/domain"
	"alisisService/internal/repository/sqlc"
)

func ToDomain(row sqlc.AdPlatform) domain.AdPlatform {
	return domain.AdPlatform{
		ID:           row.ID,
		Code:         row.Code,
		NamePlatform: row.NamePlatform,
		CreatedAt:    row.CreatedAt,
	}
}

func ToCreateAdParams(row domain.AdPlatform) sqlc.CreatePlatformParams {
	return sqlc.CreatePlatformParams{
		Code:         row.Code,
		NamePlatform: row.NamePlatform,
	}
}

func Update(id int64, dcep domain.AdPlatformPath) sqlc.UpdateAdPlatformParams {
	return sqlc.UpdateAdPlatformParams{
		NewCode:         &dcep.Code,
		NewNamePlatform: &dcep.NamePlatform,
		ID:              id,
	}
}
