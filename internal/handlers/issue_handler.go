package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"project-hive/internal/dto"
	"project-hive/internal/models"
	"project-hive/internal/service"
)

type IssueHandler struct {
	service *service.IssueService
}

func NewIssueHandler(service *service.IssueService) *IssueHandler {
	return &IssueHandler{
		service: service,
	}
}

func (h *IssueHandler) CreateIssue(c *gin.Context) {
	projectID, _ := uuid.Parse(c.Param("id"))
	var req dto.CreateIssueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	issue := models.Issue{
		ProjectID:  projectID,
		Title:      req.Title,
		Description:req.Description,
		Type:       req.Type,
		Priority:   req.Priority,
		StoryPoints:req.StoryPoints,
		SprintID:   req.SprintID,
		EpicID: req.EpicID,
	}
	if err := h.service.Create(&issue); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}
	c.JSON(http.StatusCreated, issue)
}

func (h *IssueHandler) UpdateStatus(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var req dto.UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := h.service.UpdateStatus(id, req.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func (h *IssueHandler) MoveSprint(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var req dto.MoveSprintRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := h.service.MoveToSprint(id, req.SprintID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func (h *IssueHandler) UpdateStoryPoints(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var req dto.StoryPointRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := h.service.UpdateStoryPoints(id, req.StoryPoints); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func (h *IssueHandler)GetBacklog(c *gin.Context){
	projectID,_:=uuid.Parse(c.Param("id"))
	issues,err:=h.service.GetBacklog(projectID)
	if err!=nil{
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	c.JSON(200,issues)
}

func (h *IssueHandler)MoveEpic(c *gin.Context){
	id,_:=uuid.Parse(c.Param("id"))
	var req dto.MoveEpicRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	if err:=h.service.MoveToEpic(id,req.EpicID);err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	c.Status(200)
}
