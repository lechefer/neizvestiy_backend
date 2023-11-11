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
		options.AccountId,
		options.SettlementId,
	}

	if err := r.db.SelectContext(ctx, &questsModel, query.ListQuestsSql, args...); err != nil {
		return nil, err
	}

	return mapper.QuestSliceMapFromDb(questsModel), nil
}

func (r *QuestRepository) GetQuest(ctx context.Context, accountId string, questId uuid.UUID) (entity.Quest, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return entity.Quest{}, err
	}
	defer func() { _ = tx.Rollback() }()

	var questModel models.Quest

	if err = tx.GetContext(ctx, &questModel, query.GetQuestSql, accountId, questId); err != nil {
		return entity.Quest{}, err
	}

	if questModel.Steps, err = r.getQuestSteps(ctx, tx, accountId, questId); err != nil {
		return entity.Quest{}, err
	}

	if err = tx.Commit(); err != nil {
		return entity.Quest{}, err
	}

	return mapper.QuestMapFromDb(questModel), nil
}

func (r *QuestRepository) getQuestSteps(ctx context.Context, tx *sqlx.Tx, accountId string, questId uuid.UUID) ([]models.QuestStep, error) {
	var questStepModels []models.QuestStep

	if err := tx.SelectContext(ctx, &questStepModels, query.GetQuestStepsSql, accountId, questId); err != nil {
		return nil, err
	}

	return questStepModels, nil
}

func (r *QuestRepository) StartQuest(ctx context.Context, accountId string, questId uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, query.StartQuestSql, accountId, questId)

	return err
}

func (r *QuestRepository) EndQuestStep(ctx context.Context, accountId string, questStepId uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, query.EndQuestStepSql, accountId, questStepId)

	return err
}
