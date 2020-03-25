package repository

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/types"
)

// QueryResult is store query result
type QueryResult struct {
	Result interface{}
	Error  error
}

// IssueQuery is a contract issue query
type IssueQuery interface {
	FindAllIssuesByProjectID(projectID int) <-chan QueryResult
	FindIssueByNumber(id, projectID int) <-chan QueryResult
	GetLastIssueNumberFromGivenProject(projectID int) <-chan QueryResult
}

// IssueRepository is a contract issue repository
type IssueRepository interface {
	Save(issue *domain.Issue) <-chan error
}

// MilestoneQuery is a contract milestone query
type MilestoneQuery interface {
	GetMilestones(args types.GetMilestonesInput) <-chan QueryResult
}

// MilestoneRepository is a contract milestone repository
type MilestoneRepository interface {
	Save(milestone *domain.Milestone) <-chan error
}
