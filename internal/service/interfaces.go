package service

import (
	"context"
	"github.com/google/uuid"
	"smolathon/internal/entity"
)

type accountRepo interface {
	Create(ctx context.Context, accountId string) error
}

type questRepo interface {
	ListQuests(ctx context.Context, options entity.ListQuestsOptions) ([]entity.Quest, error)
	GetQuest(ctx context.Context, accountId string, questId uuid.UUID) (entity.Quest, error)
	StartQuest(ctx context.Context, accountId string, questId uuid.UUID) error
	EndQuestStep(ctx context.Context, accountId string, questStepId uuid.UUID) error
}

type settlementRepo interface {
	ListSettlements(ctx context.Context, options entity.ListSettlementsOptions) ([]entity.Settlement, error)
}

type achievementRepo interface {
	GetAchievement(ctx context.Context, achievementId uuid.UUID, accountId string) (entity.Achievement, error)
	ListAchievements(ctx context.Context, options entity.ListAchievementsOptions) ([]entity.Achievement, error)
}

type riddleRepo interface {
	ListRiddles(ctx context.Context, options entity.ListRiddlesOptions) ([]entity.Riddle, error)
	UpdateRiddle(ctx context.Context, accountId string, riddleId uuid.UUID) error
	GetRiddle(ctx context.Context, accountId string, riddleId uuid.UUID) (entity.Riddle, error)
}
