package repository

import (
	"context"
	"smolathon/internal/entity"
	"smolathon/internal/repository/mapper"
	"smolathon/internal/repository/models"
	"smolathon/internal/repository/query"
	"smolathon/pkg/spostgres"
)

type SettlementRepository struct {
	db *spostgres.Postgres
}

func NewSettlementRepository(db *spostgres.Postgres) *SettlementRepository {
	return &SettlementRepository{
		db: db,
	}
}

func (r *SettlementRepository) ListSettlements(ctx context.Context, options entity.ListSettlementsOptions) ([]entity.Settlement, error) {
	var settlementsModels []models.Settlement

	if options.Page <= 0 {
		options.Page = 1
	}

	if options.Count <= 0 {
		options.Count = 20
	}

	offset := (options.Page - 1) * options.Count
	args := []interface{}{
		options.Query,
		offset,
		options.Count,
	}

	if err := r.db.SelectContext(ctx, &settlementsModels, query.ListSettlementsSql, args...); err != nil {
		return nil, err
	}

	return mapper.SettlementSliceMapFromDb(settlementsModels), nil
}
