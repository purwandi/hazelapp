package repository

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/types"
	uuid "github.com/satori/go.uuid"
)

type QueryResult struct {
	Result interface{}
	Error  error
}

type IssueQuery interface {
	FindAllIssuesByProjectID(projectID uuid.UUID) <-chan QueryResult
	FindIssueByIID(iid int, projectID uuid.UUID) <-chan QueryResult
	GetLastIssueNumberFromGivenProject(projectID uuid.UUID) <-chan QueryResult
}

type IssueRepository interface {
	Save(issue *domain.Issue) <-chan error
}

type MilestoneQuery interface {
	GetMilestones(args *types.GetMilestonesInput) <-chan QueryResult
}

type MilestoneRepository interface {
	Save(milestone *domain.Milestone) <-chan error
}
