package dto

import "time"

type CreateSprintRequest struct {
	Name      string    `json:"name" binding:"required"`
	Goal      string    `json:"goal"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
