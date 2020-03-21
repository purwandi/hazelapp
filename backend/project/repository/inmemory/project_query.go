package inmemory

import (
	"sort"

	"github.com/purwandi/hazelapp/project/domain"
	"github.com/purwandi/hazelapp/project/repository"
	"github.com/purwandi/hazelapp/project/storage"
	"github.com/purwandi/hazelapp/project/types"
)

type by func(p1, p2 *domain.Project) bool

func (b by) Sort(projects []domain.Project) {
	s := &projectSorter{
		projects: projects,
		by:       b,
	}

	sort.Sort(s)
}

type projectSorter struct {
	projects []domain.Project
	by       func(p1, p2 *domain.Project) bool
}

// Len is part of sort.Interface.
func (p *projectSorter) Len() int {
	return len(p.projects)
}

// Swap is part of sort.Interface.
func (p *projectSorter) Swap(i, j int) {
	p.projects[i], p.projects[j] = p.projects[j], p.projects[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (p *projectSorter) Less(i, j int) bool {
	return p.by(&p.projects[i], &p.projects[j])
}

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

		// Sort
		by(func(p1, p2 *domain.Project) bool {
			return p1.ID < p2.ID
		}).Sort(projects)

		// Filter project ownership
		for _, item := range query.Storage.ProjectMap {
			if item.OwnerID == args.OwnerID {
				projects = append(projects, item)
			}
		}

		// Calculate total query
		total = len(projects)

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
			if limit > cap(projects) {
				limit = cap(projects)
			}
			projects = projects[:limit]
		}

		if args.Last != nil {
			limit := *args.Last
			if limit > cap(projects) {
				limit = cap(projects)
			}
			offset := len(projects) - limit
			projects = projects[offset:]
		}

		// Result
		result <- repository.QueryResult{
			Result: projects,
			Total:  total,
			PageInfo: repository.PageInfo{
				StartCursor: projects[0].ID,
				EndCursor:   projects[len(projects)-1].ID,
				// Todo implement hasNextPage and endNextPage
			},
		}
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
