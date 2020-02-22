package repository

import (
	"github.com/purwandi/hazelapp/project/domain"
	uuid "github.com/satori/go.uuid"
)

type QueryResult struct {
	Result interface{}
	Error  error
}

type Repository interface {
	Save(*domain.Project) <-chan error
}

type Query interface {
	FindAll() <-chan QueryResult
	FindProjectById(uuid.UUID) <-chan QueryResult
}
