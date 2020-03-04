package inmemory

import (
	"github.com/purwandi/hazelapp/project/domain"
	"github.com/purwandi/hazelapp/project/repository"
	"github.com/purwandi/hazelapp/project/storage"
)

// ProjectRepositoryInMemory is ProjectRepository implementation in memory
type ProjectRepositoryInMemory struct {
	Storage *storage.ProjectStorage
}

// NewProjectRepositoryInMemory is to create new instance ProjectRepositoryInMemory
func NewProjectRepositoryInMemory(s *storage.ProjectStorage) repository.ProjectRepository {
	return &ProjectRepositoryInMemory{Storage: s}
}

// Save for save
func (repo *ProjectRepositoryInMemory) Save(p *domain.Project) <-chan error {
	result := make(chan error)

	go func() {
		p.ID = 1
		if len(repo.Storage.ProjectMap) > 0 {
			lastID := repo.Storage.ProjectMap[len(repo.Storage.ProjectMap)-1].ID
			p.ID = lastID + 1
		}

		repo.Storage.ProjectMap = append(repo.Storage.ProjectMap, *p)

		result <- nil
		close(result)
	}()

	return result
}
