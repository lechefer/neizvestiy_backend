package repository

import (
	"context"
	"github.com/google/uuid"
	"smolathon/internal/entity"
	"smolathon/internal/repository/mapper"
	"smolathon/internal/repository/models"
	"smolathon/internal/repository/query"
	"smolathon/pkg/spostgres"
)

type RiddleRepository struct {
	db *spostgres.Postgres
}

func NewRiddleRepository(db *spostgres.Postgres) *RiddleRepository {
	return &RiddleRepository{
		db: db,
	}
}

func (r *RiddleRepository) ListRiddles(ctx context.Context, options entity.ListRiddlesOptions) ([]entity.Riddle, error) {
	var riddlesModel []models.Riddle

	args := []interface{}{
		options.QuestStepId,
	}

	if err := r.db.SelectContext(ctx, &riddlesModel, query.ListRiddlesSql, args...); err != nil {
		return nil, err
	}

	return mapper.RiddleSliceMapFromDb(riddlesModel), nil
}

func (r *RiddleRepository) UpdateRiddle(ctx context.Context, accountId string, riddleId uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, query.UpdateRiddleSql, accountId, riddleId)
	return err
}

func (r *RiddleRepository) GetRiddle(ctx context.Context, accountId string, riddleId uuid.UUID) (entity.Riddle, error) {
	var riddleModel models.Riddle
	err := r.db.GetContext(ctx, &riddleModel, query.GetRiddleSql, accountId, riddleId)
	if err != nil {
		return entity.Riddle{}, err
	}
	return mapper.RiddleMapFromDb(riddleModel), nil
}
