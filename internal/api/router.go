package api

import (
	"net/http"
	"payment-gateway/internal/config"
	"payment-gateway/internal/middleware"
	"payment-gateway/internal/organizations"
	"payment-gateway/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(cfg config.Config, db *pgxpool.Pool) *gin.Engine {

	r := gin.Default()
	orgRepository := organizations.NewRepository(db)
	orgService := organizations.NewService(orgRepository)
	orgHandler := organizations.NewHandler(orgService)

	r.GET(
		"/health",
		middleware.AuthMiddleware(),
		func(c *gin.Context) {
			response.Success(c, http.StatusOK, "health check successful", gin.H{
				"status": "ok",
			})
		},
	)

	orgRoutes := r.Group("/organizations")
	orgRoutes.Use(middleware.AuthMiddleware())
	orgRoutes.Use(middleware.RequireRole("Admin"))
	{
		orgRoutes.POST("", orgHandler.CreateOrganization)
		orgRoutes.GET("/:id", orgHandler.GetOrganization)
		orgRoutes.GET("", orgHandler.ListOrganizations)
	}

	return r
}
