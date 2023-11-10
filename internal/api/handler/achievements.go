package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"smolathon/internal/entity"
	"smolathon/internal/service"
	"smolathon/pkg/shttp"
	"smolathon/pkg/slogger"
)

type getAchievementUri struct {
	AchievementId string `uri:"achievementId" binding:"required,uuid"`
	AccountId     string `uri:"accountId" binding:"required"`
}

type getAchievementResponse struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Steps       int       `json:"steps"`
	Description string    `json:"description"`
	Passed      int       `json:"passed"`
	IsCompleted bool      `json:"is_completed"`
}

func newGetAchievementResponse(achievement entity.Achievement) getAchievementResponse {
	response := getAchievementResponse{
		Id:          achievement.Id,
		Name:        achievement.Name,
		Icon:        achievement.Icon,
		Steps:       achievement.Steps,
		Description: achievement.Description,
		Passed:      achievement.Passed,
		IsCompleted: achievement.IsCompleted,
	}
	return response
}

// GetAchievement godoc
// @Summary     Получение достижения
// @Tags		Achievements
// @Accept      json
// @Produce     json
// @Param       AchievementId path string true "Идентификатор достижения"
// @Param       AccountId path string true "Идентификатор аккаунта"
// @Success     200 {object} shttp.ResponseWithDetails[getAchievementResponse]
// @Failure     400 {object} shttp.ResponseError "Bad request"
// @Failure     404 {object} shttp.ResponseError "Not found"
// @Failure     500 {object} shttp.ResponseError "Internal server error"
// @Router      /api/achievements/{AccountId}/{AchievementId} [get]
func GetAchievement(logger slogger.Logger, achievementService *service.AchievementService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getAchievementUri
		if err := c.ShouldBindUri(&req); err != nil {
			shttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		//TODO: при не существующем аккаунте ничего не выводить

		achievement, err := achievementService.Get(c, uuid.MustParse(req.AchievementId), req.AccountId)
		switch {
		case err == nil:
		case errors.Is(err, service.ErrNotFound):
			logger.Error(err.Error())
			shttp.ErrorResponse(c, http.StatusNotFound, service.ErrNotFound.Error())
			return
		default:
			logger.Error(err.Error())
			shttp.ErrorResponse(c, http.StatusInternalServerError, "internal  server error")
			return
		}

		response := newGetAchievementResponse(achievement)
		shttp.OkResponseWithResult(c, response)
	}
}

type listAchievementsRequest struct {
	AccountId string `form:"account_id" binding:"required"`
}

type listAchievementResponse []listAchievementResponseElement

type listAchievementResponseElement struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Icon        string    `json:"icon"`
	Steps       int       `json:"steps"`
	Description string    `json:"description"`
	Passed      int       `json:"passed"`
	IsCompleted bool      `json:"is_completed"`
}

func newListAchievementsResponse(achievements []entity.Achievement) listAchievementResponse {
	var response = make(listAchievementResponse, 0, len(achievements))
	for _, achievement := range achievements {
		response = append(response, listAchievementResponseElement{
			Id:          achievement.Id,
			Name:        achievement.Name,
			Icon:        achievement.Icon,
			Steps:       achievement.Steps,
			Description: achievement.Description,
			Passed:      achievement.Passed,
			IsCompleted: achievement.IsCompleted,
		})
	}
	return response
}

// ListAchievements godoc
// @Summary     Получение достижений
// @Tags		Achievements
// @Accept      json
// @Produce     json
// @Param       QueryParams query listAchievementsRequest true "Параметры выборки"
// @Success     200 {object} shttp.ResponseWithDetails[listAchievementResponse]
// @Failure     400 {object} shttp.ResponseError "Bad request"
// @Failure     500 {object} shttp.ResponseError "Internal server error"
// @Router      /api/achievements/list [get]
func ListAchievements(logger slogger.Logger, achievementService *service.AchievementService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req listAchievementsRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			shttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}
		options := entity.ListAchievementsOptions{
			AccountId: req.AccountId,
		}

		achievements, err := achievementService.List(c, options)
		switch {
		case err == nil:
		default:
			logger.Error(err.Error())
			shttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
			return
		}

		response := newListAchievementsResponse(achievements)
		shttp.OkResponseWithResult(c, response)
	}
}
