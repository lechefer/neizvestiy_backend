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

type AchievementRepository struct {
	db *spostgres.Postgres
}

func NewAchievementRepository(db *spostgres.Postgres) *AchievementRepository {
	return &AchievementRepository{
		db: db,
	}
}

func (r *AchievementRepository) GetAchievement(ctx context.Context, achievementId uuid.UUID, accountId string) (entity.Achievement, error) {
	var achievementModel models.Achievement
	if err := r.db.GetContext(ctx, &achievementModel, query.GetAchievementSql, accountId, achievementId); err != nil {
		return entity.Achievement{}, err
	}

	return mapper.AchievementMapFromDb(achievementModel), nil
}

func (r *AchievementRepository) ListAchievements(ctx context.Context, options entity.ListAchievementsOptions) ([]entity.Achievement, error) {
	var achievementModel []models.Achievement
	args := []interface{}{
		options.AccountId,
	}

	if err := r.db.SelectContext(ctx, &achievementModel, query.ListAchievementsSql, args...); err != nil {
		return nil, err
	}

	return mapper.AchievementSliceFromDb(achievementModel), nil
}
