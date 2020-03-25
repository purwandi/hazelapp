package services

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/types"
)

// IssueService is to expose issue service
type IssueService struct {
	query      repository.IssueQuery
	repository repository.IssueRepository
}

// NewIssueService is to create IssueService instance
func NewIssueService(query repository.IssueQuery, repo repository.IssueRepository) *IssueService {
	return &IssueService{
		query:      query,
		repository: repo,
	}
}

// GetLastIssueNumberFromGivenProject is to get last issue number from given project
func (s *IssueService) GetLastIssueNumberFromGivenProject(id int) (int, error) {
	// Process
	result := <-s.query.GetLastIssueNumberFromGivenProject(id)
	if result.Error != nil {
		return 0, result.Error
	}

	// Response
	return result.Result.(int), nil
}

// CreateIssue is to create an issue
func (s *IssueService) CreateIssue(args types.CreateIssueInput) (domain.Issue, error) {
	// Process
	issue, err := domain.CreateIssue(s, args)
	if err != nil {
		return domain.Issue{}, err
	}

	// Persist
	err = <-s.repository.Save(issue)
	if err != nil {
		return domain.Issue{}, err
	}

	// Response
	return *issue, nil
}
