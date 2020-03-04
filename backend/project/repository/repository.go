package repository

import (
	"github.com/purwandi/hazelapp/project/domain"
	"github.com/purwandi/hazelapp/project/types"
)

// QueryResult is to wrap query result
type QueryResult struct {
	Result interface{}
	Error  error
}

// ProjectRepository is an interface for project repository
type ProjectRepository interface {
	Save(*domain.Project) <-chan error
}

// ProjectQuery is an interface for project query
type ProjectQuery interface {
	GetProjects(types.GetProjectsInput) <-chan QueryResult
	FindProject(types.FindProjectInput) <-chan QueryResult
}
