package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"project-hive/internal/models"
)

type SprintRepository struct {
	DB *gorm.DB
}

func NewSprintRepository(db *gorm.DB) *SprintRepository {
	return &SprintRepository{
		DB: db,
	}
}

func (r *SprintRepository) Create(sprint *models.Sprint) error {
	return r.DB.Create(sprint).Error
}

func (r *SprintRepository) FindByID(id uuid.UUID) (*models.Sprint, error) {
	var sprint models.Sprint

	err := r.DB.First(&sprint, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &sprint, nil
}

func (r *SprintRepository) FindByProject(projectID uuid.UUID) ([]models.Sprint, error) {
	var sprints []models.Sprint

	err := r.DB.Where("project_id = ?", projectID).
		Find(&sprints).Error

	return sprints, err
}

func (r *SprintRepository) Update(sprint *models.Sprint) error {
	return r.DB.Save(sprint).Error
}

func (r *SprintRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Sprint{}, "id = ?", id).Error
}
