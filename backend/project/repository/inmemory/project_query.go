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
	total := 0

	go func() {
		projects := []domain.Project{}

		if args.First == nil && args.Last == nil {
			result <- repository.QueryResult{
				Error: repository.ProjectQueryError{Code: repository.ProjectQueryErrorMissingPaginationBoundaries},
			}
		}

		if args.First != nil && args.Last != nil {
			result <- repository.QueryResult{
				Error: repository.ProjectQueryError{Code: repository.ProjectQueryErrorPagination},
			}
		}

		// Filter project ownership
		for _, item := range query.Storage.ProjectMap {
			if item.OwnerID == args.OwnerID {
				projects = append(projects, item)
			}
		}

		// Sort
		by(func(p1, p2 *domain.Project) bool {
			return p1.ID < p2.ID
		}).Sort(projects)

		// Calculate total query
		total = len(projects)
		if total == 0 {
			result <- repository.QueryResult{
				Result: projects,
			}
		}

		// Cursor
		if args.After != nil {
			var offset int
			for index, item := range projects {
				if item.ID > *args.After {
					offset = index
					break
				}
			}
			projects = projects[offset:]
		}

		if args.Before != nil {
			var offset int
			for index, item := range projects {
				if item.ID == *args.Before {
					offset = index
				}
			}
			projects = projects[:offset]
		}

		// Get n first items
		if args.First != nil {
			limit := *args.First
			if limit > len(projects) {
				limit = len(projects)
			}
			projects = projects[:limit]
		}

		if args.Last != nil {
			limit := *args.Last
			if limit > len(projects) {
				limit = len(projects)
			}
			offset := len(projects) - limit
			projects = projects[offset:]
		}

		// Result
		res := repository.QueryResult{
			Result: projects,
			Total:  total,
		}

		if len(projects) > 0 {
			res.PageInfo = repository.PageInfo{
				StartCursor: projects[0].ID,
				EndCursor:   projects[len(projects)-1].ID,
				// Todo implement hasNextPage and endNextPage
			}
		}

		result <- res
		close(result)
	}()

	return result
}

// FindProjectByID is to find a project
func (query *ProjectQueryInMemory) FindProjectByID(id int) <-chan repository.QueryResult {
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
