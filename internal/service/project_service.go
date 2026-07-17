package service

import (
	"github.com/google/uuid"

	"project-hive/internal/models"
	"project-hive/internal/repository"
)

type ProjectService struct {
	repo *repository.ProjectRepository
}

func NewProjectService(repo *repository.ProjectRepository) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (s *ProjectService) Create(project *models.Project) error {
	project.ID = uuid.New()

	return s.repo.Create(project)
}

func (s *ProjectService) GetAll() ([]models.Project, error) {
	return s.repo.FindAll()
}

func (s *ProjectService) GetByID(id uuid.UUID) (*models.Project, error) {
	return s.repo.FindByID(id)
}

func (s *ProjectService) Update(project *models.Project) error {
	return s.repo.Update(project)
}

func (s *ProjectService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}
