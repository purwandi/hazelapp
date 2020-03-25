package inmemory

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/storage"
	"github.com/purwandi/hazelapp/issue/types"
)

// IssueQueryInMemory is to IssueQuery implementation in memory
type IssueQueryInMemory struct {
	Storage *storage.IssueStorage
}

// NewIssueQueryInMemory is to create IssueQueryInMemory instance
func NewIssueQueryInMemory(s *storage.IssueStorage) repository.IssueQuery {
	return &IssueQueryInMemory{Storage: s}
}

// GetLastIssueNumberFromGivenProject is to get last issue number from given project
func (query *IssueQueryInMemory) GetLastIssueNumberFromGivenProject(id int) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		issue := domain.Issue{}
		for _, item := range query.Storage.IssueMap {
			if item.ProjectID == id && item.Number > issue.Number {
				issue = item
			}
		}

		if (domain.Issue{}) == issue {
			result <- repository.QueryResult{Result: 0}
		}

		result <- repository.QueryResult{Result: issue.Number}
		close(result)
	}()

	return result
}

// GetIssues to get all issues
func (query *IssueQueryInMemory) GetIssues(args types.GetIssuesInput) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		issues := []domain.Issue{}
		for _, issue := range query.Storage.IssueMap {
			if issue.ProjectID == args.ProjectID {
				issues = append(issues, issue)
			}
		}

		result <- repository.QueryResult{Result: issues}
		close(result)
	}()

	return result
}

// GetIssueByNumber to get an issue
func (query *IssueQueryInMemory) GetIssueByNumber(args types.GetIssueInput) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		issue := domain.Issue{}
		for _, item := range query.Storage.IssueMap {
			if item.Number == args.Number && item.ProjectID == args.ProjectID {
				issue = item
			}
		}

		result <- repository.QueryResult{Result: issue}
		close(result)
	}()

	return result
}
