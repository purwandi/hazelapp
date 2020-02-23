package repository

import (
	"github.com/purwandi/hazelapp/issue/domain"
	uuid "github.com/satori/go.uuid"
)

type QueryResult struct {
	Result interface{}
	Error  error
}

type IssueQuery interface {
	FindAllIssuesByProjectID(uuid.UUID) <-chan QueryResult
	GetLastIssueNumberFromGivenProject(uuid.UUID) <-chan QueryResult
}

type IssueRepository interface {
	Save(*domain.Issue) <-chan error
}
