package routes

import (
	"github.com/gin-gonic/gin"

	"project-hive/internal/handlers"
)

func RegisterRoutes(
	router *gin.Engine,
	projectHandler *handlers.ProjectHandler,
	sprintHandler *handlers.SprintHandler,
	issueHandler *handlers.IssueHandler,
) {

	api := router.Group("/api/v1")

	{
		projects := api.Group("/projects")

		{
			projects.POST("", projectHandler.CreateProject)
			projects.GET("", projectHandler.GetProjects)
			projects.GET("/:id", projectHandler.GetProject)
			projects.DELETE("/:id", projectHandler.DeleteProject)

			projects.POST("/:id/sprints", sprintHandler.CreateSprint)
			projects.POST("/:id/issues", issueHandler.CreateIssue)
		}

		sprints := api.Group("/sprints")

		{
			sprints.POST("/:id/start", sprintHandler.StartSprint)
			sprints.POST("/:id/end", sprintHandler.FinishSprint)
		}

		issues := api.Group("/issues")

		{
			issues.PATCH("/:id/status", issueHandler.UpdateStatus)
			issues.PATCH("/:id/story-points", issueHandler.UpdateStoryPoints)
			issues.PATCH("/:id/sprint", issueHandler.MoveSprint)
		}
	}
}
