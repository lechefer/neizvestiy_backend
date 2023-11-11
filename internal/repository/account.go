package repository

import (
	"context"
	"smolathon/internal/repository/query"
	"smolathon/pkg/spostgres"
)

type AccountRepository struct {
	db *spostgres.Postgres
}

func NewAccountRepository(db *spostgres.Postgres) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Create(ctx context.Context, accountId string) error {
	_, err := r.db.ExecContext(ctx, query.CreateAccountSql, accountId)
	return err
}
