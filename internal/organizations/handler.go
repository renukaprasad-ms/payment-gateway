package organizations

import (
	"errors"
	"net/http"
	"payment-gateway/internal/response"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateOrganization(c *gin.Context) {
	var req CreateOrganizationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	org, err := h.service.CreateOrganization(c.Request.Context(), req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "organization created successfully", org)
}

func (h *Handler) GetOrganization(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error(c, http.StatusBadRequest, "organization id is required")
		return
	}

	org, err := h.service.GetOrganization(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			response.Error(c, http.StatusNotFound, "organization not found")
			return
		}
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "organization fetched successfully", org)
}

func (h *Handler) ListOrganizations(c *gin.Context) {
	org, err := h.service.GetAllOrganization(c.Request.Context())
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusOK, "organization fetched successfully", org)
}
