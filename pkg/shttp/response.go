package shttp

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

type NoResult interface{}

type ResponseOk ResponseWithDetails[NoResult]

type ResponseError ResponseWithDetails[NoResult]

type ResponseWithDetails[T any] struct {
	Ok          bool   `json:"ok"`
	Result      T      `json:"result,omitempty"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`
}

func writeResponse[T any](c *gin.Context, r ResponseWithDetails[T]) {
	c.Writer.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(c.Writer).Encode(r); err != nil {
		// if error without result
		if reflect.TypeOf(r.Result) != reflect.TypeOf(NoResult(nil)) {
			_ = c.Error(err)
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		// if error with result
		_ = c.Error(err)
		ErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}
}

func OkResponse(c *gin.Context) {
	resp := ResponseWithDetails[NoResult]{
		Ok: true,
	}

	writeResponse(c, resp)
}

func OkResponseWithResult[T any](c *gin.Context, result T) {
	resp := ResponseWithDetails[T]{
		Ok:     true,
		Result: result,
	}

	writeResponse(c, resp)
}

func ErrorResponse(c *gin.Context, code int, msg string) {
	resp := ResponseError{
		Ok:          false,
		ErrorCode:   code,
		Description: msg,
	}

	writeResponse[NoResult](c, ResponseWithDetails[NoResult](resp))
}
