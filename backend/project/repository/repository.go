package repository

import (
	"github.com/purwandi/hazelapp/project/domain"
	uuid "github.com/satori/go.uuid"
)

type QueryResult struct {
	Result interface{}
	Error  error
}

type ProjectRepository interface {
	Save(*domain.Project) <-chan error
}

type ProjectQuery interface {
	FindAll() <-chan QueryResult
	FindProjectByID(uuid.UUID) <-chan QueryResult
}
