package service

import (
	"github.com/google/uuid"

	"project-hive/internal/models"
	"project-hive/internal/repository"
)

type EpicService struct {
	repo *repository.EpicRepository
}

func NewEpicService(repo *repository.EpicRepository)*EpicService{
	return &EpicService{
		repo:repo,
	}
}

func (s *EpicService) Create(epic *models.Epic)error{
	epic.ID = uuid.New()
	return s.repo.Create(epic)
}

func (s *EpicService)Get(id uuid.UUID)(*models.Epic,error){
	return s.repo.FindByID(id)
}

func (s *EpicService)GetProjectEpics(projectID uuid.UUID)([]models.Epic,error){
	return s.repo.FindByProject(projectID)
}

func (s *EpicService)Update(epic *models.Epic)error{
	return s.repo.Update(epic)
}

func (s *EpicService)Delete(id uuid.UUID)error{
	return s.repo.Delete(id)
}

func (s *IssueService)GetBacklog(projectID uuid.UUID)([]models.Issue,error){
	return s.repo.FindBacklog(projectID)
}
