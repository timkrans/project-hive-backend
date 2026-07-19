package dto

import "github.com/google/uuid"

type BacklogResponse struct {
	ProjectID uuid.UUID `json:"project_id"`
	Issues []IssueResponse `json:"issues"`
}
