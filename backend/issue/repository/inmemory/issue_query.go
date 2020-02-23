package inmemory

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/storage"
	uuid "github.com/satori/go.uuid"
)

type IssueQueryInMemory struct {
	Storage *storage.IssueStorage
}

func NewIssueQueryInMemory(s *storage.IssueStorage) repository.IssueQuery {
	return &IssueQueryInMemory{Storage: s}
}

func (query *IssueQueryInMemory) GetLastIssueNumberFromGivenProject(id uuid.UUID) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		issue := domain.Issue{}
		for _, item := range query.Storage.IssueMap {
			if item.ProjectID == id && item.IID > issue.IID {
				issue = item
			}
		}

		if (domain.Issue{}) == issue {
			result <- repository.QueryResult{Result: 0}
		}

		result <- repository.QueryResult{Result: issue.IID}
		close(result)
	}()

	return result
}

func (query *IssueQueryInMemory) FindAllIssuesByProjectID(id uuid.UUID) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		issues := []domain.Issue{}
		for _, issue := range query.Storage.IssueMap {
			if issue.ProjectID == id {
				issues = append(issues, issue)
			}
		}

		result <- repository.QueryResult{Result: issues}
		close(result)
	}()

	return result
}

func (query *IssueQueryInMemory) FindIssueByIID(id int, projectID uuid.UUID) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		issue := domain.Issue{}
		for _, item := range query.Storage.IssueMap {
			if item.IID == id && item.ProjectID == projectID {
				issue = item
			}
		}

		result <- repository.QueryResult{Result: issue}
		close(result)
	}()

	return result
}
