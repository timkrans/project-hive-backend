package dto

import "github.com/google/uuid"

type CreateEpicRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateEpicRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type EpicResponse struct {
	ID uuid.UUID `json:"id"`

	ProjectID uuid.UUID `json:"project_id"`

	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`

	IssueCount int `json:"issue_count"`
}
