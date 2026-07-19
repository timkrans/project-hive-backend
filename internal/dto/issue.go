package dto

import (
	"github.com/google/uuid"

	"project-hive/internal/models"
)

type CreateIssueRequest struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description"`
	Type models.IssueType `json:"type"`
	Priority models.Priority `json:"priority"`
	StoryPoints int `json:"story_points"`
	//optional
	EpicID *uuid.UUID `json:"epic_id"`
	// optional
	SprintID *uuid.UUID `json:"sprint_id"`
}

type StoryPointRequest struct {
	StoryPoints int `json:"story_points"`
}

type UpdateIssueRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Type models.IssueType `json:"type"`
	Priority models.Priority `json:"priority"`
	StoryPoints int `json:"story_points"`
}

type UpdateStatusRequest struct {
	Status models.IssueStatus `json:"status" binding:"required"`
}

type UpdateStoryPointsRequest struct {
	StoryPoints int `json:"story_points"`
}

type MoveSprintRequest struct {
	//null moves issue back to backlog
	SprintID *uuid.UUID `json:"sprint_id"`
}

type MoveEpicRequest struct {
	//null removes epic association
	EpicID *uuid.UUID `json:"epic_id"`
}

type IssueResponse struct {
	ID uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
	EpicID *uuid.UUID `json:"epic_id"`
	SprintID *uuid.UUID `json:"sprint_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Type models.IssueType `json:"type"`
	Status models.IssueStatus `json:"status"`
	Priority models.Priority `json:"priority"`
	StoryPoints int `json:"story_points"`
	Position int `json:"position"`
}
