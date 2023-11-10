package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"smolathon/internal/entity"
	"smolathon/internal/service"
	"smolathon/pkg/shttp"
	"smolathon/pkg/slogger"
)

type searchSettlementsQuery struct {
	Query string `form:"query"`
}

type searchSettlementsResponse []searchSettlementsResponseElement

type searchSettlementsResponseElement struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func newSearchSettlementsResponse(settlements []entity.Settlement) searchSettlementsResponse {
	var response = make(searchSettlementsResponse, 0, len(settlements))
	for _, settlement := range settlements {
		response = append(response, searchSettlementsResponseElement{
			Id:   settlement.Id,
			Name: settlement.Name,
		})
	}
	return response
}

// SearchSettlements godoc
// @Summary     Поиск города
// @Tags		Settlements
// @Accept      json
// @Produce     json
// @Param       QueryParams query searchSettlementsQuery true "Параметры поиска"
// @Success     200 {object} shttp.ResponseWithDetails[searchSettlementsResponse]
// @Failure     400 {object} shttp.ResponseError "Bad request"
// @Failure     500 {object} shttp.ResponseError "Internal server error"
// @Router      /api/settlements/search [get]
func SearchSettlements(logger slogger.Logger, settlementService *service.SettlementService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req searchSettlementsQuery
		if err := c.ShouldBindQuery(&req); err != nil {
			shttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		options := entity.ListSettlementsOptions{
			Query: req.Query,
			Page:  1,
			Count: 20,
		}

		settlements, err := settlementService.List(c, options)
		switch {
		case err == nil:
		default:
			logger.Error(err.Error())
			shttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
			return
		}

		response := newSearchSettlementsResponse(settlements)
		shttp.OkResponseWithResult(c, response)
	}
}
