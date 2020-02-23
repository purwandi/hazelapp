package services

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	uuid "github.com/satori/go.uuid"
)

type IssueService struct {
	query      repository.IssueQuery
	repository repository.IssueRepository
}

func NewIssueService(query repository.IssueQuery, repo repository.IssueRepository) *IssueService {
	return &IssueService{
		query:      query,
		repository: repo,
	}
}

func (s *IssueService) GetLastIssueNumberFromGivenProject(id uuid.UUID) (int, error) {
	// Process
	result := <-s.query.GetLastIssueNumberFromGivenProject(id)
	if result.Error != nil {
		return 0, result.Error
	}

	// Response
	return result.Result.(int), nil
}

func (s *IssueService) CreateIssue(title, body string, projectID uuid.UUID) (domain.Issue, error) {
	// Process
	issue, err := domain.CreateIssue(s, title, body, projectID)
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

func (s *IssueService) UpdateIssue(iid int, title, body string, projectID uuid.UUID) (domain.Issue, error) {
	// Process
	result := <-s.query.FindIssueByIID(iid, projectID)
	if result.Error != nil {
		return domain.Issue{}, result.Error
	}

	issue := result.Result.(*domain.Issue)
	issue.ChangeTitle(title)
	issue.ChangeBody(body)

	// Persist
	if err := <-s.repository.Save(issue); err != nil {
		return domain.Issue{}, err
	}

	// Response
	return *issue, nil
}
