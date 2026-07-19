package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"project-hive/internal/dto"
	"project-hive/internal/models"
	"project-hive/internal/service"
)

type EpicHandler struct {
	service *service.EpicService
}

func (h *IssueHandler) GetIssue(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid issue id",
		})
		return
	}
	issue, err := h.service.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(
		http.StatusOK,
		issue,
	)
}

func NewEpicHandler(service *service.EpicService)*EpicHandler{
	return &EpicHandler{
		service:service,
	}
}

func (h *EpicHandler)CreateEpic(c *gin.Context){
	projectID,_:=uuid.Parse(c.Param("id"))
	var req dto.CreateEpicRequest
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	epic:=models.Epic{
		ProjectID:projectID,
		Name:req.Name,
		Description:req.Description,
	}
	if err:=h.service.Create(&epic);err!=nil{

		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated,epic)
}

func (h *EpicHandler)GetProjectEpics(c *gin.Context){
	projectID,_:=uuid.Parse(c.Param("id"))
	epics,err:=h.service.GetProjectEpics(projectID)
	if err!=nil{

		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	c.JSON(200,epics)
}

func (h *EpicHandler)GetEpic(c *gin.Context){
	id,_:=uuid.Parse(c.Param("id"))
	epic,err:=h.service.Get(id)
	if err!=nil{
		c.JSON(404,gin.H{"error":err.Error()})
		return
	}
	c.JSON(200,epic)
}

func (h *EpicHandler)DeleteEpic(c *gin.Context){
	id,_:=uuid.Parse(c.Param("id"))
	if err:=h.service.Delete(id);err!=nil{
		c.JSON(400,gin.H{"error":err.Error()})
		return
	}
	c.Status(204)
}
