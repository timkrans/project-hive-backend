package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateSprintRequest struct {
	Name string `json:"name" binding:"required"`
	Goal string `json:"goal"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
}

type UpdateSprintRequest struct {
	Name string `json:"name"`
	Goal string `json:"goal"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
}

type SprintResponse struct {
	ID uuid.UUID `json:"id"`
	ProjectID uuid.UUID `json:"project_id"`
	Name string `json:"name"`
	Goal string `json:"goal"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
	Active bool `json:"active"`
	IssueCount int `json:"issue_count"`
}

