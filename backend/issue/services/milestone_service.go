package services

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/types"
)

// MilestoneService is expose milestone service
type MilestoneService struct {
	query      repository.MilestoneQuery
	repository repository.MilestoneRepository
}

// NewMilestoneService is to create new MilestoneService instance
func NewMilestoneService(query repository.MilestoneQuery, repo repository.MilestoneRepository) *MilestoneService {
	return &MilestoneService{
		query:      query,
		repository: repo,
	}
}

// CreateMilestone is to create milestone
func (s *MilestoneService) CreateMilestone(args types.CreateMilestoneInput) (domain.Milestone, error) {
	milestone, err := domain.CreateMilestone(args)
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

// GetMilestones is to get milestones
func (s *MilestoneService) GetMilestones(args *types.GetMilestonesInput) ([]domain.Milestone, error) {
	// Validate
	if args == nil {
		return []domain.Milestone{}, ServiceError{ServiceErrorArgsIsBlank}
	}

	// Process
	result := <-s.query.GetMilestones(args)
	if result.Error != nil {
		return []domain.Milestone{}, result.Error
	}

	// Response
	return result.Result.([]domain.Milestone), nil
}
