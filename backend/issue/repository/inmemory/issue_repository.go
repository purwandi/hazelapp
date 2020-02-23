package inmemory

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/storage"
)

type IssueRepositoryInMemory struct {
	Storage *storage.IssueStorage
}

func NewIssueRepositoryInMemory(s *storage.IssueStorage) repository.IssueRepository {
	return &IssueRepositoryInMemory{Storage: s}
}

func (repo *IssueRepositoryInMemory) Save(issue *domain.Issue) <-chan error {
	result := make(chan error)

	go func() {
		repo.Storage.IssueMap = append(repo.Storage.IssueMap, *issue)

		result <- nil
		close(result)
	}()

	return result
}
