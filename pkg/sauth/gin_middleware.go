package sauth

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"smolathon/pkg/shttp"
)

func RequireRole(enforcer casbin.IEnforcer, requireRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		accountId := GetAccountId(c)

		if hasAccess, _ := enforcer.HasRoleForUser(accountId.String(), requireRole); !hasAccess {
			shttp.ErrorResponse(c, http.StatusForbidden, "access denied")
			c.Abort()
			return
		}

		c.Next()
	}
}
