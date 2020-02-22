package inmemory

import (
	"github.com/purwandi/hazelapp/project/domain"
	"github.com/purwandi/hazelapp/project/repository"
	"github.com/purwandi/hazelapp/project/storage"
)

type ProjectRepositoryInMemory struct {
	Storage *storage.ProjectStorage
}

func NewProjectRepositoryInMemory(s *storage.ProjectStorage) repository.Repository {
	return &ProjectRepositoryInMemory{Storage: s}
}

func (repo *ProjectRepositoryInMemory) Save(p *domain.Project) <-chan error {
	result := make(chan error)

	go func() {
		repo.Storage.ProjectMap = append(repo.Storage.ProjectMap, *p)

		result <- nil
		close(result)
	}()

	return result
}
