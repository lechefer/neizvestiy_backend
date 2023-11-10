package sauth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"smolathon/pkg/shttp"
	"strings"
)

const (
	AccountIdKey = "accountId"
	UsernameKey  = "username"
)

type Authorize struct {
	AccountId uuid.UUID `json:"account_id"`
	Username  string    `json:"username"`
}

type AccessTokenParser interface {
	Parse(tokenStr string) (Authorize, error)
}

func BearerAuth(atp AccessTokenParser) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if !strings.HasPrefix(token, "Bearer ") {
			shttp.ErrorResponse(c, http.StatusForbidden, "invalid token")
			c.Abort()
			return
		}

		authorize, err := atp.Parse(strings.TrimPrefix(token, "Bearer "))
		if err != nil {
			shttp.ErrorResponse(c, http.StatusForbidden, "invalid token")
			c.Abort()
			return
		}

		c.Set(AccountIdKey, authorize.AccountId)
		c.Set(UsernameKey, authorize.Username)

		c.Next()
	}
}

func GetAccountId(ctx context.Context) uuid.UUID {
	val, ok := ctx.Value(AccountIdKey).(uuid.UUID)
	if !ok {
		return uuid.Nil
	}
	return val
}

func GetParamUUID(c *gin.Context, key string) uuid.UUID {
	id, _ := uuid.Parse(c.Params.ByName(key))
	return id
}
