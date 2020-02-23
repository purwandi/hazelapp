package services

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	uuid "github.com/satori/go.uuid"
)

type MilestoneService struct {
	query      repository.MilestoneQuery
	repository repository.MilestoneRepository
}

func NewMilestoneService(query repository.MilestoneQuery, repo repository.MilestoneRepository) *MilestoneService {
	return &MilestoneService{
		query:      query,
		repository: repo,
	}
}

func (s *MilestoneService) CreateMilestone(name, description string, projectID uuid.UUID) (domain.Milestone, error) {
	// Process
	milestone, err := domain.CreateMilestone(name, description, projectID)
	if err != nil {
		return domain.Milestone{}, err
	}

	// Persist
	if err := <-s.repository.Save(milestone); err != nil {
		return domain.Milestone{}, err
	}

	// Response
	return *milestone, nil
}

func (s *MilestoneService) FindAllMilestonesByProjectID(projectID uuid.UUID) ([]domain.Milestone, error) {
	// Process
	result := <-s.query.FindAllMilestonesByProjectID(projectID)
	if result.Error != nil {
		return []domain.Milestone{}, result.Error
	}

	// Response
	return result.Result.([]domain.Milestone), nil
}
