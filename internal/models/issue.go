package models

import (
	"time"

	"github.com/google/uuid"
)

type Issue struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey"`

	ProjectID uuid.UUID  `gorm:"type:char(36);not null"`
	SprintID  *uuid.UUID `gorm:"type:char(36)"`

	Title       string `gorm:"not null"`
	Description string

	StoryPoints int

	Type     IssueType
	Status   IssueStatus
	Priority Priority

	Position int

	CreatedAt time.Time
	UpdatedAt time.Time
}
