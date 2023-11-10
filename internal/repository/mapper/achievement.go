package mapper

import (
	"smolathon/internal/entity"
	"smolathon/internal/repository/models"
)

func AchievementMapFromDb(achievement models.Achievement) entity.Achievement {
	return entity.Achievement{
		Id:          achievement.Id,
		Name:        achievement.Name,
		Icon:        achievement.Icon,
		Steps:       achievement.Steps,
		Description: achievement.Description,
		Passed:      achievement.Passed,
		IsCompleted: achievement.IsCompleted,
	}
}

func AchievementSliceFromDb(achievementModels []models.Achievement) []entity.Achievement {
	var achievements = make([]entity.Achievement, 0, len(achievementModels))
	for _, achievement := range achievementModels {
		achievements = append(achievements, AchievementMapFromDb(achievement))
	}
	return achievements
}
