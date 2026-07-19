package models

import (
	"time"

	"github.com/google/uuid"
)

type Issue struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey"`
	ProjectID uuid.UUID `gorm:"type:char(36);not null"`
	//null = issue is not part of an epic
	EpicID *uuid.UUID `gorm:"type:char(36)"`
	//null = issue is in backlog
	SprintID *uuid.UUID `gorm:"type:char(36)"`
	Title string `gorm:"not null"`
	Description string
	Type IssueType
	Status IssueStatus
	Priority Priority
	StoryPoints int
	Position int
	CreatedAt time.Time
	UpdatedAt time.Time
}
