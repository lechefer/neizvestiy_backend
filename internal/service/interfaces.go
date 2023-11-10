package service

import (
	"context"
	"github.com/google/uuid"
	"smolathon/internal/entity"
)

type questRepo interface {
	ListQuests(ctx context.Context, options entity.ListQuestsOptions) ([]entity.Quest, error)
	GetQuest(ctx context.Context, questId uuid.UUID) (entity.Quest, error)
}

type settlementRepo interface {
	ListSettlements(ctx context.Context, options entity.ListSettlementsOptions) ([]entity.Settlement, error)
}

type achievementRepo interface {
	GetAchievement(ctx context.Context, achievementId uuid.UUID, accountId string) (entity.Achievement, error)
	ListAchievements(ctx context.Context, options entity.ListAchievementsOptions) ([]entity.Achievement, error)
}
