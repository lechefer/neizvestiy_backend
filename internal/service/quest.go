package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"smolathon/internal/entity"
	"strings"
)

type QuestService struct {
	imageService *ImageService
	questRepo    questRepo
}

func NewQuestService(imageService *ImageService, questRepo questRepo) *QuestService {
	return &QuestService{
		imageService: imageService,
		questRepo:    questRepo,
	}
}

func (qs *QuestService) List(ctx context.Context, options entity.ListQuestsOptions) ([]entity.Quest, error) {
	quests, err := qs.questRepo.ListQuests(ctx, options)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil
	default:
		return nil, err
	}

	for i, quest := range quests {
		quests[i].Preview, err = qs.imageService.GetQuestPreview(ctx, quest.Id)
		if err != nil {
			return nil, err
		}
	}

	return quests, nil
}

func (qs *QuestService) Get(ctx context.Context, accountId string, questId uuid.UUID) (entity.Quest, error) {
	quest, err := qs.questRepo.GetQuest(ctx, accountId, questId)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return entity.Quest{}, ErrNotFound
	default:
		return entity.Quest{}, err
	}

	quest.Preview, err = qs.imageService.GetQuestPreview(ctx, quest.Id)
	if err != nil {
		return entity.Quest{}, err
	}

	for i, step := range quest.Steps {
		quest.Steps[i].Images, err = qs.imageService.GetQuestStepImages(ctx, quest.Id, step.Id)
		if err != nil {
			return entity.Quest{}, err
		}
	}

	return quest, nil
}

func (qs *QuestService) Start(ctx context.Context, accountId string, questId uuid.UUID) (entity.Quest, error) {
	err := qs.questRepo.StartQuest(ctx, accountId, questId)
	switch {
	case err == nil:
	case strings.Contains(err.Error(), "duplicate key"):
		return entity.Quest{}, ErrAlreadyExists
	default:
		return entity.Quest{}, err
	}

	quest, err := qs.Get(ctx, accountId, questId)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return entity.Quest{}, ErrNotFound
	default:
		return entity.Quest{}, err
	}

	return quest, nil
}

func (qs *QuestService) EndStep(ctx context.Context, accountId string, questStepId uuid.UUID) (entity.Quest, error) {
	err := qs.questRepo.EndQuestStep(ctx, accountId, questStepId)
	switch {
	case err == nil:
	default:
		return entity.Quest{}, err
	}

	quest, err := qs.Get(ctx, accountId, questStepId)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return entity.Quest{}, ErrNotFound
	default:
		return entity.Quest{}, err
	}

	return quest, nil
}
