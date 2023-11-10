package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"smolathon/internal/entity"
)

type AchievementService struct {
	achievementRepo achievementRepo
}

func NewAchievementService(achievementRepo achievementRepo) *AchievementService {
	return &AchievementService{
		achievementRepo: achievementRepo,
	}
}

func (as *AchievementService) Get(ctx context.Context, achievementId uuid.UUID, accountId string) (entity.Achievement, error) {
	achievement, err := as.achievementRepo.GetAchievement(ctx, achievementId, accountId)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return entity.Achievement{}, ErrNotFound
	default:
		return entity.Achievement{}, err
	}
	return achievement, nil
}

func (as *AchievementService) List(ctx context.Context, options entity.ListAchievementsOptions) ([]entity.Achievement, error) {
	achievements, err := as.achievementRepo.ListAchievements(ctx, options)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil
	default:
		return nil, err
	}
	return achievements, nil
}
