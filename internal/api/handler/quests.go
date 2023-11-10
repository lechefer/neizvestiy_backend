package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"net/http"
	"smolathon/internal/entity"
	"smolathon/internal/service"
	"smolathon/pkg/shttp"
	"smolathon/pkg/slogger"
)

type image struct {
	Sizes sizes `json:"sizes"`
}

type sizes struct {
	M size `json:"m"`
	X size `json:"x"`
	O size `json:"o"`
}

type size struct {
	Size string `json:"size"`
	Url  string `json:"url"`
}

type listQuestsRequest struct {
	SettlementId string `form:"settlement_id" binding:"required,uuid"`
}

type listQuestsResponse []listQuestsResponseElement

type listQuestsResponseElement struct {
	Id           uuid.UUID        `json:"id"`
	SettlementId uuid.UUID        `json:"settlement_id"`
	Preview      image            `json:"preview"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Type         entity.QuestType `json:"type"`
	Duration     int              `json:"duration"`
	Reward       decimal.Decimal  `json:"reward"`
}

func newListQuestsResponse(quests []entity.Quest) listQuestsResponse {
	var response = make(listQuestsResponse, 0, len(quests))
	for _, quest := range quests {
		response = append(response, listQuestsResponseElement{
			Id:           quest.Id,
			SettlementId: quest.SettlementId,
			Preview: image{Sizes: sizes{
				M: size{
					Size: string(quest.Preview.Sizes.M.Size),
					Url:  quest.Preview.Sizes.M.Url,
				},
				X: size{
					Size: string(quest.Preview.Sizes.X.Size),
					Url:  quest.Preview.Sizes.X.Url,
				},
				O: size{
					Size: string(quest.Preview.Sizes.O.Size),
					Url:  quest.Preview.Sizes.O.Url,
				},
			}},
			Name:        quest.Name,
			Description: quest.Description,
			Type:        quest.Type,
			Duration:    int(quest.AvgDuration.Minutes()),
			Reward:      quest.Reward,
		})
	}
	return response
}

// ListQuests godoc
// @Summary     Получение квестов
// @Tags		Quests
// @Accept      json
// @Produce     json
// @Param       QueryParams query listQuestsRequest true "Параметры выборки"
// @Success     200 {object} shttp.ResponseWithDetails[listQuestsResponse]
// @Failure     400 {object} shttp.ResponseError "Bad request"
// @Failure     500 {object} shttp.ResponseError "Internal server error"
// @Router      /api/quests/list [get]
func ListQuests(logger slogger.Logger, questService *service.QuestService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req listQuestsRequest
		if err := c.ShouldBindQuery(&req); err != nil {
			shttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		options := entity.ListQuestsOptions{
			SettlementId: uuid.MustParse(req.SettlementId),
		}

		quests, err := questService.List(c, options)
		switch {
		case err == nil:
		default:
			logger.Error(err.Error())
			shttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
			return
		}

		response := newListQuestsResponse(quests)
		shttp.OkResponseWithResult(c, response)
	}
}

type listQuestsUri struct {
	QuestId string `uri:"questId" binding:"required,uuid"`
}

type getQuestResponse struct {
	Id           uuid.UUID              `json:"id"`
	SettlementId uuid.UUID              `json:"settlement_id"`
	Preview      image                  `json:"preview"`
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	Type         entity.QuestType       `json:"type"`
	Duration     int                    `json:"duration"`
	Reward       decimal.Decimal        `json:"reward"`
	Steps        []getQuestResponseStep `json:"steps"`
}

type getQuestResponseStep struct {
	Id        uuid.UUID                      `json:"id"`
	QuestId   uuid.UUID                      `json:"quest_id"`
	Order     int                            `json:"order"`
	Images    []image                        `json:"images"`
	Name      string                         `json:"name"`
	PlaceType string                         `json:"place_type"`
	Address   string                         `json:"address"`
	Phone     string                         `json:"phone"`
	Email     string                         `json:"email"`
	Website   string                         `json:"website"`
	Schedule  []getQuestResponseStepSchedule `json:"schedule"`
	Latitude  float64                        `json:"latitude"`
	Longitude float64                        `json:"longitude"`
}

type getQuestResponseStepSchedule struct {
	WeekDay string `json:"week_day"`
	FromTo  string `json:"from_to"`
}

func newGetQuestResponse(quest entity.Quest) getQuestResponse {
	response := getQuestResponse{
		Id:           quest.Id,
		SettlementId: quest.SettlementId,
		Preview: image{Sizes: sizes{
			M: size{
				Size: string(quest.Preview.Sizes.M.Size),
				Url:  quest.Preview.Sizes.M.Url,
			},
			X: size{
				Size: string(quest.Preview.Sizes.X.Size),
				Url:  quest.Preview.Sizes.X.Url,
			},
			O: size{
				Size: string(quest.Preview.Sizes.O.Size),
				Url:  quest.Preview.Sizes.O.Url,
			},
		}},
		Name:        quest.Name,
		Description: quest.Description,
		Type:        quest.Type,
		Duration:    int(quest.AvgDuration.Minutes()),
		Reward:      quest.Reward,
	}

	response.Steps = make([]getQuestResponseStep, 0, len(quest.Steps))
	for _, questStep := range quest.Steps {
		step := getQuestResponseStep{
			Id:        questStep.Id,
			QuestId:   questStep.QuestId,
			Order:     questStep.Order,
			Name:      questStep.Name,
			PlaceType: questStep.PlaceType,
			Address:   questStep.Address,
			Phone:     questStep.Phone,
			Email:     questStep.Email,
			Website:   questStep.Website,
			Latitude:  questStep.Latitude,
			Longitude: questStep.Longitude,
		}

		step.Schedule = make([]getQuestResponseStepSchedule, 0, len(questStep.Schedule))
		for _, schedule := range questStep.Schedule {
			step.Schedule = append(step.Schedule, getQuestResponseStepSchedule{
				WeekDay: string(schedule.WeekDay),
				FromTo:  schedule.FromTo,
			})
		}

		step.Images = make([]image, 0, len(questStep.Images))
		for _, questStepImage := range questStep.Images {
			step.Images = append(step.Images, image{Sizes: sizes{
				M: size{
					Size: string(questStepImage.Sizes.M.Size),
					Url:  questStepImage.Sizes.M.Url,
				},
				X: size{
					Size: string(questStepImage.Sizes.X.Size),
					Url:  questStepImage.Sizes.X.Url,
				},
				O: size{
					Size: string(questStepImage.Sizes.O.Size),
					Url:  questStepImage.Sizes.O.Url,
				},
			}})
		}

		response.Steps = append(response.Steps, step)
	}

	return response
}

// GetQuest godoc
// @Summary     Получение информации о квесте
// @Tags		Quests
// @Accept      json
// @Produce     json
// @Param       questId path string true "Идентификатор квеста"
// @Success     200 {object} shttp.ResponseWithDetails[getQuestResponse]
// @Failure     400 {object} shttp.ResponseError "Bad request"
// @Failure     404 {object} shttp.ResponseError "Not found"
// @Failure     500 {object} shttp.ResponseError "Internal server error"
// @Router      /api/quests/{questId} [get]
func GetQuest(logger slogger.Logger, questService *service.QuestService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req listQuestsUri
		if err := c.ShouldBindUri(&req); err != nil {
			shttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		quest, err := questService.Get(c, uuid.MustParse(req.QuestId))
		switch {
		case err == nil:
		case errors.Is(err, service.ErrNotFound):
			logger.Error(err.Error())
			shttp.ErrorResponse(c, http.StatusNotFound, service.ErrNotFound.Error())
			return
		default:
			logger.Error(err.Error())
			shttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
			return
		}

		response := newGetQuestResponse(quest)
		shttp.OkResponseWithResult(c, response)
	}
}
