package repository

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/types"
)

// PageInfo is describe pagination
type PageInfo struct {
	StartCursor     int
	EndCursor       int
	HasNextPage     *bool
	HasPreviousPage *bool
}

// QueryResult is store query result
type QueryResult struct {
	Result   interface{}
	Total    int
	Error    error
	PageInfo PageInfo
}

// IssueQuery is a contract issue query
type IssueQuery interface {
	GetIssues(types.GetIssuesInput) <-chan QueryResult
	GetIssueByNumber(types.GetIssueInput) <-chan QueryResult
	GetLastIssueNumberFromGivenProject(projectID int) <-chan QueryResult
}

// IssueRepository is a contract issue repository
type IssueRepository interface {
	Save(issue *domain.Issue) <-chan error
}

// MilestoneQuery is a contract milestone query
type MilestoneQuery interface {
	FindMilestoneByID(int) <-chan QueryResult
	GetMilestones(types.GetMilestonesInput) <-chan QueryResult
}

// MilestoneRepository is a contract milestone repository
type MilestoneRepository interface {
	Save(milestone *domain.Milestone) <-chan error
}
