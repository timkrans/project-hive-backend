package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"project-hive/internal/dto"
	"project-hive/internal/models"
	"project-hive/internal/service"
)

type SprintHandler struct {
	service *service.SprintService
}

func NewSprintHandler(service *service.SprintService) *SprintHandler {
	return &SprintHandler{
		service: service,
	}
}

func (h *SprintHandler) CreateSprint(c *gin.Context) {

	projectID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid project id",
		})
		return
	}

	var req dto.CreateSprintRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	sprint := models.Sprint{
		ProjectID: projectID,
		Name:      req.Name,
		Goal:      req.Goal,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}

	if err := h.service.Create(&sprint); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, sprint)
}

func (h *SprintHandler) StartSprint(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("id"))

	if err := h.service.Start(id); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.Status(http.StatusOK)
}

func (h *SprintHandler) FinishSprint(c *gin.Context) {

	id, _ := uuid.Parse(c.Param("id"))

	if err := h.service.Finish(id); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.Status(http.StatusOK)
}
