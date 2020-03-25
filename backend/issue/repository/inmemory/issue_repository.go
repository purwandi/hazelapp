package inmemory

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/storage"
)

// IssueRepositoryInMemory is IssueRepository implementation in memory
type IssueRepositoryInMemory struct {
	Storage *storage.IssueStorage
}

// NewIssueRepositoryInMemory to create IssueRepositoryInMemory instance
func NewIssueRepositoryInMemory(s *storage.IssueStorage) repository.IssueRepository {
	return &IssueRepositoryInMemory{Storage: s}
}

// Save for save
func (repo *IssueRepositoryInMemory) Save(issue *domain.Issue) <-chan error {
	result := make(chan error)

	go func() {
		issue.ID = 1
		if len(repo.Storage.IssueMap) > 0 {
			lastID := repo.Storage.IssueMap[len(repo.Storage.IssueMap)-1].ID
			issue.ID = lastID + 1
		}

		repo.Storage.IssueMap = append(repo.Storage.IssueMap, *issue)

		result <- nil
		close(result)
	}()

	return result
}
