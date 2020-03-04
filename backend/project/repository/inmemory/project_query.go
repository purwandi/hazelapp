package inmemory

import (
	"github.com/purwandi/hazelapp/project/domain"
	"github.com/purwandi/hazelapp/project/repository"
	"github.com/purwandi/hazelapp/project/storage"
	"github.com/purwandi/hazelapp/project/types"
)

// ProjectQueryInMemory is ProjectQuery implementation in memory
type ProjectQueryInMemory struct {
	Storage *storage.ProjectStorage
}

// NewProjectQueryInMemory is to create NewProjectQueryInMemory instance
func NewProjectQueryInMemory(s *storage.ProjectStorage) repository.ProjectQuery {
	return &ProjectQueryInMemory{Storage: s}
}

// GetProjects is to get all projects
func (query *ProjectQueryInMemory) GetProjects(args types.GetProjectsInput) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		result <- repository.QueryResult{Result: query.Storage.ProjectMap}
		close(result)
	}()

	return result
}

// FindProject is to find a project
func (query *ProjectQueryInMemory) FindProject(args types.FindProjectInput) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		project := domain.Project{}
		for _, p := range query.Storage.ProjectMap {
			if p.Name == args.Name && p.OwnerID == args.OwnerID {
				project = p
			}
		}

		result <- repository.QueryResult{Result: project}
		close(result)
	}()

	return result
}
