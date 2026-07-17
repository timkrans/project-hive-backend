package main

import (
	"github.com/gin-gonic/gin"

	"project-hive/internal/database"
	"project-hive/internal/handlers"
	"project-hive/internal/repository"
	"project-hive/internal/routes"
	"project-hive/internal/service"
)

func main() {

	database.ConnectDatabase()

	projectRepo := repository.NewProjectRepository(database.DB)
	sprintRepo := repository.NewSprintRepository(database.DB)
	issueRepo := repository.NewIssueRepository(database.DB)

	projectService := service.NewProjectService(projectRepo)
	sprintService := service.NewSprintService(sprintRepo)
	issueService := service.NewIssueService(issueRepo)

	projectHandler := handlers.NewProjectHandler(projectService)
	sprintHandler := handlers.NewSprintHandler(sprintService)
	issueHandler := handlers.NewIssueHandler(issueService)

	router := gin.Default()

	routes.RegisterRoutes(
		router,
		projectHandler,
		sprintHandler,
		issueHandler,
	)

	router.Run(":8080")
}
