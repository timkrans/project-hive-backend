package dto

import "github.com/google/uuid"

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Key         string `json:"key" binding:"required"`
	Description string `json:"description"`
}

type UpdateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Key         string    `json:"key"`
	Description string    `json:"description"`
	CreatedAt string `json:"created_at"`
}
