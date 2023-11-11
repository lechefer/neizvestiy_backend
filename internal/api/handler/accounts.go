package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"smolathon/internal/service"
	"smolathon/pkg/shttp"
	"smolathon/pkg/slogger"
)

type accountCreateRequest struct {
	AccountId string `json:"account_id"`
}

// AccountCreate godoc
// @Summary     Создание аккаунта
// @Tags		Accounts
// @Accept      json
// @Produce     json
// @Param       data body accountCreateRequest true "Данные аккаунта"
// @Success     200 {object} shttp.ResponseWithDetails[string]
// @Failure     400 {object} shttp.ResponseError "Bad request"
// @Failure     409 {object} shttp.ResponseError "Already exists"
// @Failure     500 {object} shttp.ResponseError "Internal server error"
// @Router      /api/accounts/create [post]
func AccountCreate(logger slogger.Logger, accountService *service.AccountService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req accountCreateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			shttp.ErrorResponse(c, http.StatusBadRequest, "bad body")
			return
		}

		err := accountService.Create(c, req.AccountId)
		switch {
		case err == nil:
		case errors.Is(err, service.ErrAlreadyExists):
			logger.Error(err.Error())
			shttp.ErrorResponse(c, http.StatusConflict, service.ErrAlreadyExists.Error())
			return
		default:
			logger.Error(err.Error())
			shttp.ErrorResponse(c, http.StatusInternalServerError, "internal server error")
			return
		}

		shttp.OkResponse(c)
	}
}
