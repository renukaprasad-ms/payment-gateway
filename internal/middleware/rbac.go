package middleware

import (
	"net/http"
	"payment-gateway/internal/response"

	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {

	return func(c *gin.Context) {

		roles, exists := c.Get("roles")
		if !exists {
			response.AbortError(c, http.StatusForbidden, "no roles")
			return
		}

		roleList := roles.([]string)

		for _, r := range roleList {
			if r == role {
				c.Next()
				return
			}
		}

		response.AbortError(c, http.StatusForbidden, "access denied")
	}
}
