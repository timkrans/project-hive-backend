package models

import (
	"time"

	"github.com/google/uuid"
)

type Epic struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey"`
	ProjectID uuid.UUID `gorm:"type:char(36);not null"`
	Name        string `gorm:"not null"`
	Description string
	Status string `gorm:"default:'open'"`
	Issues []Issue `gorm:"foreignKey:EpicID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
