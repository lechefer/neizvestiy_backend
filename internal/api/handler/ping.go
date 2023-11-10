package handler

import (
	"github.com/gin-gonic/gin"
	"smolathon/pkg/shttp"
)

// Ping godoc
// @Summary     Пинг сервиса
// @Tags		Служебные
// @Accept      json
// @Produce     json
// @Success     200 {object} shttp.ResponseWithDetails[string]
// @Router      /api/ping [get]
func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		shttp.OkResponseWithResult(c, "pong")
	}
}
