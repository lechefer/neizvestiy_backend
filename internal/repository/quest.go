package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"smolathon/internal/entity"
	"smolathon/internal/repository/mapper"
	"smolathon/internal/repository/models"
	"smolathon/internal/repository/query"
	"smolathon/pkg/spostgres"
)

type QuestRepository struct {
	db *spostgres.Postgres
}

func NewQuestRepository(db *spostgres.Postgres) *QuestRepository {
	return &QuestRepository{
		db: db,
	}
}

func (r *QuestRepository) ListQuests(ctx context.Context, options entity.ListQuestsOptions) ([]entity.Quest, error) {
	var questsModel []models.Quest

	args := []interface{}{
		options.SettlementId,
	}

	if err := r.db.SelectContext(ctx, &questsModel, query.ListQuestsSql, args...); err != nil {
		return nil, err
	}

	return mapper.QuestSliceMapFromDb(questsModel), nil
}

func (r *QuestRepository) GetQuest(ctx context.Context, questId uuid.UUID) (entity.Quest, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return entity.Quest{}, err
	}
	defer func() { _ = tx.Rollback() }()

	var questModel models.Quest

	if err = tx.GetContext(ctx, &questModel, query.GetQuestSql, questId); err != nil {
		return entity.Quest{}, err
	}

	if questModel.Steps, err = r.getQuestSteps(ctx, tx, questId); err != nil {
		return entity.Quest{}, err
	}

	if err = tx.Commit(); err != nil {
		return entity.Quest{}, err
	}

	return mapper.QuestMapFromDb(questModel), nil
}

func (r *QuestRepository) getQuestSteps(ctx context.Context, tx *sqlx.Tx, questId uuid.UUID) ([]models.QuestStep, error) {
	var questStepModels []models.QuestStep

	if err := tx.SelectContext(ctx, &questStepModels, query.GetQuestStepsSql, questId); err != nil {
		return nil, err
	}

	return questStepModels, nil
}
