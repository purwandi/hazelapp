package services

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/types"
)

// IssueCollectionResult is to wrap issue collection result
type IssueCollectionResult struct {
	Issues   []domain.Issue
	Total    int
	PageInfo repository.PageInfo
}

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

	// Return
	return result.Result.(int), nil
}

// GetIssues is to get issues
func (s *IssueService) GetIssues(args types.GetIssuesInput) (IssueCollectionResult, error) {
	// Process
	result := <-s.query.GetIssues(args)
	if result.Error != nil {
		return IssueCollectionResult{}, result.Error
	}

	// Return
	return IssueCollectionResult{
		Issues:   result.Result.([]domain.Issue),
		Total:    result.Total,
		PageInfo: result.PageInfo,
	}, nil
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

	// Return
	return *issue, nil
}
