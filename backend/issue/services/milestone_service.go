package services

import (
	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/types"
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

func (s *MilestoneService) CreateMilestone(args *types.CreateMilestoneInput) (domain.Milestone, error) {
	// Process
	projectID, err := uuid.FromString(args.ProjectID)
	if err != nil {
		return domain.Milestone{}, err
	}

	milestone, err := domain.CreateMilestone(args.Name, helpers.StringValue(args.Description), projectID)
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

func (s *MilestoneService) FindAllMilestones(args *types.GetMilestoneQueryInput) ([]domain.Milestone, error) {
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
