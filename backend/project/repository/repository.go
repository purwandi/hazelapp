package repository

import (
	"github.com/purwandi/hazelapp/project/domain"
	"github.com/purwandi/hazelapp/project/types"
)

// PageInfo is describe pagination
type PageInfo struct {
	StartCursor     int
	EndCursor       int
	HasNextPage     *bool
	HasPreviousPage *bool
}

// QueryResult is to wrap query result
type QueryResult struct {
	Result   interface{}
	Total    int
	Error    error
	PageInfo PageInfo
}

// ProjectRepository is an interface for project repository
type ProjectRepository interface {
	Save(*domain.Project) <-chan error
}

// ProjectQuery is an interface for project query
type ProjectQuery interface {
	GetProjects(types.GetProjectsInput) <-chan QueryResult
	FindProjectByID(int) <-chan QueryResult
	FindProject(types.FindProjectInput) <-chan QueryResult
}
