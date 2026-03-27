package middleware

import (
	"log"
	"net/http"
	"payment-gateway/internal/response"

	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {

	return func(c *gin.Context) {

		roles, exists := c.Get("roles")
		if !exists {
			log.Printf("rbac deny: missing roles required=%s method=%s path=%s", role, c.Request.Method, c.Request.URL.Path)
			response.AbortError(c, http.StatusForbidden, "no roles")
			return
		}

		roleList, ok := roles.([]string)
		if !ok {
			log.Printf("rbac deny: invalid roles type=%T required=%s method=%s path=%s", roles, role, c.Request.Method, c.Request.URL.Path)
			response.AbortError(c, http.StatusForbidden, "invalid roles")
			return
		}

		log.Printf("rbac check: required=%s user_id=%s roles=%v method=%s path=%s", role, c.GetString("user_id"), roleList, c.Request.Method, c.Request.URL.Path)

		for _, r := range roleList {
			if r == role {
				log.Printf("rbac allow: required=%s matched_role=%s user_id=%s method=%s path=%s", role, r, c.GetString("user_id"), c.Request.Method, c.Request.URL.Path)
				c.Next()
				return
			}
		}

		log.Printf("rbac deny: access denied required=%s user_id=%s roles=%v method=%s path=%s", role, c.GetString("user_id"), roleList, c.Request.Method, c.Request.URL.Path)
		response.AbortError(c, http.StatusForbidden, "access denied")
	}
}
