package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"project-hive/internal/models"
)

type ProjectRepository struct {
	DB *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		DB: db,
	}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return r.DB.Create(project).Error
}

func (r *ProjectRepository) FindAll() ([]models.Project, error) {
	var projects []models.Project

	err := r.DB.Find(&projects).Error

	return projects, err
}

func (r *ProjectRepository) FindByID(id uuid.UUID) (*models.Project, error) {
	var project models.Project

	err := r.DB.First(&project, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *ProjectRepository) Update(project *models.Project) error {
	return r.DB.Save(project).Error
}

func (r *ProjectRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Project{}, "id = ?", id).Error
}
