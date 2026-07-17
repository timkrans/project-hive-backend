package models

import (
	"time"

	"github.com/google/uuid"
)

type Sprint struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey"`

	ProjectID uuid.UUID `gorm:"type:char(36);not null"`

	Name string `gorm:"not null"`
	Goal string

	StartDate time.Time
	EndDate   time.Time

	Active bool

	Issues []Issue `gorm:"foreignKey:SprintID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
