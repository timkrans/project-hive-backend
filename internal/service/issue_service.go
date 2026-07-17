package service

import (
	"errors"

	"github.com/google/uuid"

	"project-hive/internal/models"
	"project-hive/internal/repository"
)

type IssueService struct {
	repo *repository.IssueRepository
}

func NewIssueService(repo *repository.IssueRepository) *IssueService {
	return &IssueService{
		repo: repo,
	}
}

func (s *IssueService) Create(issue *models.Issue) error {
	issue.ID = uuid.New()

	if issue.Status == "" {
		issue.Status = models.Todo
	}

	return s.repo.Create(issue)
}

func (s *IssueService) Get(id uuid.UUID) (*models.Issue, error) {
	return s.repo.FindByID(id)
}

func (s *IssueService) GetProjectIssues(projectID uuid.UUID) ([]models.Issue, error) {
	return s.repo.FindByProject(projectID)
}

func (s *IssueService) Update(issue *models.Issue) error {
	return s.repo.Update(issue)
}

func (s *IssueService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *IssueService) UpdateStatus(id uuid.UUID, status models.IssueStatus) error {
	issue, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	switch status {
	case models.Todo,
		models.InProgress,
		models.Review,
		models.Done,
		models.Blocked:
	default:
		return errors.New("invalid status")
	}

	issue.Status = status

	return s.repo.Update(issue)
}

func (s *IssueService) UpdateStoryPoints(id uuid.UUID, points int) error {
	issue, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	issue.StoryPoints = points

	return s.repo.Update(issue)
}

func (s *IssueService) MoveToSprint(issueID uuid.UUID, sprintID uuid.UUID) error {
	issue, err := s.repo.FindByID(issueID)
	if err != nil {
		return err
	}

	issue.SprintID = &sprintID

	return s.repo.Update(issue)
}
