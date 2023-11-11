package mapper

import (
	"smolathon/internal/entity"
	"smolathon/internal/repository/models"
)

func RiddleSliceMapFromDb(riddlesModels []models.Riddle) []entity.Riddle {
	var riddles = make([]entity.Riddle, 0, len(riddlesModels))
	for _, riddle := range riddlesModels {
		riddles = append(riddles, RiddleMapFromDb(riddle))
	}
	return riddles
}

func RiddleMapFromDb(riddle models.Riddle) entity.Riddle {
	return entity.Riddle{
		Id:          riddle.Id,
		QuestStepId: riddle.QuestStepId,
		Name:        riddle.Name,
		Description: riddle.Description,
		Status:      riddle.Status,
		Letters:     riddle.Letter,
	}
}
