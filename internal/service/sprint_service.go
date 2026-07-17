package service

import (
	"github.com/google/uuid"

	"project-hive/internal/models"
	"project-hive/internal/repository"
)

type SprintService struct {
	repo *repository.SprintRepository
}

func NewSprintService(repo *repository.SprintRepository) *SprintService {
	return &SprintService{
		repo: repo,
	}
}

func (s *SprintService) Create(sprint *models.Sprint) error {
	sprint.ID = uuid.New()

	return s.repo.Create(sprint)
}

func (s *SprintService) Get(id uuid.UUID) (*models.Sprint, error) {
	return s.repo.FindByID(id)
}

func (s *SprintService) GetProjectSprints(projectID uuid.UUID) ([]models.Sprint, error) {
	return s.repo.FindByProject(projectID)
}

func (s *SprintService) Update(sprint *models.Sprint) error {
	return s.repo.Update(sprint)
}

func (s *SprintService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *SprintService) Start(id uuid.UUID) error {
	sprint, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	sprint.Active = true

	return s.repo.Update(sprint)
}

func (s *SprintService) Finish(id uuid.UUID) error {
	sprint, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	sprint.Active = false

	return s.repo.Update(sprint)
}
