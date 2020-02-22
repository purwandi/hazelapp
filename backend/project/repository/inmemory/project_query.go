package inmemory

import (
	"github.com/purwandi/hazelapp/project/domain"
	"github.com/purwandi/hazelapp/project/repository"
	"github.com/purwandi/hazelapp/project/storage"
	uuid "github.com/satori/go.uuid"
)

type ProjectQueryInMemory struct {
	Storage *storage.ProjectStorage
}

func NewProjectQueryInMemory(s *storage.ProjectStorage) repository.Query {
	return &ProjectQueryInMemory{Storage: s}
}

func (query *ProjectQueryInMemory) FindAll() <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		projects := []domain.Project{}
		for _, p := range query.Storage.ProjectMap {
			projects = append(projects, p)
		}

		result <- repository.QueryResult{Result: projects}
		close(result)
	}()

	return result
}

func (query *ProjectQueryInMemory) FindProjectById(id uuid.UUID) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		project := domain.Project{}
		for _, p := range query.Storage.ProjectMap {
			if p.ID == id {
				project = p
			}
		}

		result <- repository.QueryResult{Result: project}
		close(result)
	}()

	return result
}
