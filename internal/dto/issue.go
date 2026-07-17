package dto

import (
	"github.com/google/uuid"

	"project-hive/internal/models"
)

type CreateIssueRequest struct {
	Title       string             `json:"title" binding:"required"`
	Description string             `json:"description"`
	Type        models.IssueType   `json:"type"`
	Priority    models.Priority    `json:"priority"`
	StoryPoints int                `json:"story_points"`
	SprintID    *uuid.UUID         `json:"sprint_id"`
}

type UpdateStatusRequest struct {
	Status models.IssueStatus `json:"status"`
}

type StoryPointRequest struct {
	StoryPoints int `json:"story_points"`
}

type MoveSprintRequest struct {
	SprintID uuid.UUID `json:"sprint_id"`
}
