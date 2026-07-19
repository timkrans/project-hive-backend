package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"project-hive/internal/models"
)

type EpicRepository struct {
	DB *gorm.DB
}

func NewEpicRepository(db *gorm.DB) *EpicRepository {
	return &EpicRepository{
		DB: db,
	}
}

func (r *EpicRepository) Create(epic *models.Epic) error {
	return r.DB.Create(epic).Error
}

func (r *EpicRepository) FindByID(id uuid.UUID) (*models.Epic,error) {
	var epic models.Epic
	err := r.DB.
		Preload("Issues").
		First(&epic,"id = ?",id).
		Error
	if err != nil {
		return nil,err
	}
	return &epic,nil
}

func (r *EpicRepository) FindByProject(projectID uuid.UUID)([]models.Epic,error){
	var epics []models.Epic
	err := r.DB.
		Where("project_id = ?",projectID).
		Find(&epics).
		Error
	return epics,err
}

func (r *EpicRepository) Update(epic *models.Epic) error {
	return r.DB.Save(epic).Error
}

func (r *EpicRepository) Delete(id uuid.UUID) error {
	return r.DB.
		Delete(&models.Epic{},"id = ?",id).
		Error
}
