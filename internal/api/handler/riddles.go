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

type listRiddlesRequest struct {
	AccountId   string `form:"account_id" binding:"required"`
	QuestStepId string `form:"quest_step_id" binding:"required,uuid"`
}

type listRiddlesResponse []listRiddlesResponseElement

type listRiddlesResponseElement struct {
	Id          uuid.UUID `json:"id"`
	QuestStepId uuid.UUID `json:"quest_step_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Letters     string    `json:"letters"`
}

func newListRiddlesResponse(riddles []entity.Riddle) listRiddlesResponse {
	var response = make(listRiddlesResponse, 0, len(riddles))
	for _, riddle := range riddles {
		response = append(response, listRiddlesResponseElement{
			Id:          riddle.Id,
			QuestStepId: riddle.QuestStepId,
			Name:        riddle.Name,
			Description: riddle.Description,
			Status:      riddle.Status,
			Letters:     riddle.Letters,
		})
	}
	return response
}

// ListRiddles godoc
// @Summary     Получение квестов
// @Tags		Riddles
// @Accept      json
// @Produce     json
// @Param       QueryParams query listRiddlesRequest true "Параметры выборки"
// @Success     200 {object} shttp.ResponseWithDetails[listRiddlesResponse]
// @Failure     400 {object} shttp.ResponseError "Bad request"
// @Failure     500 {object} shttp.ResponseError "Internal server error"
// @Router      /api/riddles/list [get]
func ListRiddles(logger slogger.Logger, riddleService *service.RiddleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req listRiddlesRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			shttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		options := entity.ListRiddlesOptions{
			QuestStepId: uuid.MustParse(req.QuestStepId),
		}
		quests, err := riddleService.List(c, options)
		switch {
		case err == nil:
		default:
			logger.Error(err.Error())
			shttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
			return
		}
		response := newListRiddlesResponse(quests)
		shttp.OkResponseWithResult(c, response)
	}
}

type updateRiddleUri struct {
	AccountId string `uri:"accountId" binding:"required"`
	RiddleId  string `uri:"riddleId" binding:"required,uuid"`
}

type updateRiddleResponse struct {
	Id          uuid.UUID `json:"id"`
	QuestStepId uuid.UUID `json:"quest_step_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Letters     string    `json:"letters"`
}

func newUpdateRiddleResponse(riddle entity.Riddle) updateRiddleResponse {
	response := updateRiddleResponse{
		Id:          riddle.Id,
		QuestStepId: riddle.QuestStepId,
		Name:        riddle.Name,
		Description: riddle.Description,
		Status:      riddle.Status,
		Letters:     riddle.Letters,
	}
	return response
}

// UpdateRiddle godoc
// @Summary     Обновление загадки
// @Tags		Riddles
// @Accept      json
// @Produce     json
// @Param       accountId path string true "Идентификатор пользователя"
// @Param       riddleId path string true "Идентификатор загадки"
// @Success     200 {object} shttp.ResponseWithDetails[updateRiddleResponse]
// @Failure     400 {object} shttp.ResponseError "Bad request"
// @Failure     409 {object} shttp.ResponseError "Already exists"
// @Failure     500 {object} shttp.ResponseError "Internal server error"
// @Router      /api/riddles/{accountId}/{riddleId} [post]
func UpdateRiddle(logger slogger.Logger, riddleService *service.RiddleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req updateRiddleUri
		if err := c.ShouldBindUri(&req); err != nil {
			shttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		riddle, err := riddleService.Update(c, req.AccountId, uuid.MustParse(req.RiddleId))
		switch {
		case err == nil:
		case errors.Is(err, service.ErrNotFound):
			logger.Error(err.Error())
			shttp.ErrorResponse(c, http.StatusNotFound, service.ErrNotFound.Error())
			return
		case errors.Is(err, service.ErrAlreadyExists):
			logger.Error(err.Error())
			shttp.ErrorResponse(c, http.StatusConflict, service.ErrAlreadyExists.Error())
			return
		default:
			logger.Error(err.Error())
			shttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
			return
		}
		response := newUpdateRiddleResponse(riddle)
		shttp.OkResponseWithResult(c, response)
	}
}
