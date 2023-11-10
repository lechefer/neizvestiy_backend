package service

import (
	"context"
	"database/sql"
	"errors"
	"smolathon/internal/entity"
)

type SettlementService struct {
	settlementRepo settlementRepo
}

func NewSettlementService(settlementRepo settlementRepo) *SettlementService {
	return &SettlementService{
		settlementRepo: settlementRepo,
	}
}

func (ss *SettlementService) List(ctx context.Context, options entity.ListSettlementsOptions) ([]entity.Settlement, error) {
	settlements, err := ss.settlementRepo.ListSettlements(ctx, options)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil
	default:
		return nil, err
	}

	return settlements, nil
}
