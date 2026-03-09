package api

import (
	"net/http"
	"payment-gateway/internal/config"
	"payment-gateway/internal/middleware"
	"payment-gateway/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(cfg config.Config, db *pgxpool.Pool) *gin.Engine {

	r := gin.Default()

	r.GET(
		"/health",
		middleware.AuthMiddleware(),
		func(c *gin.Context) {
			response.Success(c, http.StatusOK, "health check successful", gin.H{
				"status": "ok",
			})
		},
	)

	return r
}
