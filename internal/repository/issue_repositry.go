package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"project-hive/internal/models"
)

type IssueRepository struct {
	DB *gorm.DB
}

func NewIssueRepository(db *gorm.DB) *IssueRepository {
	return &IssueRepository{
		DB: db,
	}
}

func (r *IssueRepository) Create(issue *models.Issue) error {
	return r.DB.Create(issue).Error
}

func (r *IssueRepository) FindByID(id uuid.UUID) (*models.Issue, error) {
	var issue models.Issue

	err := r.DB.First(&issue, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &issue, nil
}

func (r *IssueRepository) FindByProject(projectID uuid.UUID) ([]models.Issue, error) {
	var issues []models.Issue
	err := r.DB.
		Where("project_id = ?", projectID).
		Find(&issues).Error

	return issues, err
}

func (r *IssueRepository) Update(issue *models.Issue) error {
	return r.DB.Save(issue).Error
}

func (r *IssueRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&models.Issue{}, "id = ?", id).Error
}

func (r *IssueRepository) FindBacklog(projectID uuid.UUID)([]models.Issue,error){
	var issues []models.Issue
	err:=r.DB.
		Where(
			"project_id = ? AND sprint_id IS NULL",
			projectID,
		).
		Order("position ASC").
		Find(&issues).
		Error
	return issues,err
}
