package models

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Name        string    `gorm:"not null"`
	Key         string    `gorm:"unique;not null"`
	Description string

	Sprints []Sprint `gorm:"foreignKey:ProjectID"`
	Issues  []Issue  `gorm:"foreignKey:ProjectID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
