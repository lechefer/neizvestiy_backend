package mapper

import (
	"encoding/json"
	"smolathon/internal/entity"
	"smolathon/internal/repository/models"
)

func QuestSliceMapFromDb(questsModels []models.Quest) []entity.Quest {
	var quests = make([]entity.Quest, 0, len(questsModels))
	for _, quest := range questsModels {
		quests = append(quests, QuestMapFromDb(quest))
	}
	return quests
}

func QuestMapFromDb(quest models.Quest) entity.Quest {
	return entity.Quest{
		Id:           quest.Id,
		SettlementId: quest.SettlementId,
		Name:         quest.Name,
		Description:  quest.Description,
		Type:         entity.QuestType(quest.Type),
		AvgDuration:  quest.AvgDuration,
		Reward:       quest.Reward,
		Steps:        QuestStepSliceMapFromDb(quest.Steps),
	}
}

func QuestStepSliceMapFromDb(questStepsModels []models.QuestStep) []entity.QuestStep {
	var questSteps = make([]entity.QuestStep, 0, len(questStepsModels))
	for _, questStep := range questStepsModels {
		questSteps = append(questSteps, QuestStepMapFromDb(questStep))
	}
	return questSteps
}

func QuestStepMapFromDb(questStep models.QuestStep) entity.QuestStep {
	var schedule []entity.Schedule
	_ = json.Unmarshal([]byte(questStep.Schedule), &schedule)

	return entity.QuestStep{
		Id:        questStep.Id,
		QuestId:   questStep.QuestId,
		Order:     questStep.Order,
		Name:      questStep.Name,
		PlaceType: questStep.PlaceType,
		Address:   questStep.Address,
		Phone:     questStep.Phone,
		Email:     questStep.Email,
		Website:   questStep.Website,
		Schedule:  schedule,
		Latitude:  questStep.Location.P.X,
		Longitude: questStep.Location.P.Y,
	}
}
