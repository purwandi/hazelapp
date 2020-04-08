package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
)

// IssueResolver resolve ...
type IssueResolver struct {
	Field    domain.Issue
	Resolver *Resolver
}

// ID resolve ...
func (i *IssueResolver) ID() graphql.ID {
	return graphql.ID(helpers.EncryptID("issue", i.Field.ID))
}

// Project resolve ...
func (i *IssueResolver) Project() (*ProjectResolver, error) {
	project, err := i.Resolver.ProjectService.FindProjectByID(i.Field.ProjectID)
	if err != nil {
		return nil, err
	}

	return &ProjectResolver{
		Field:    project,
		Resolver: i.Resolver,
	}, nil
}

// Milestone resolve ...
func (i *IssueResolver) Milestone() (*MilestoneResolver, error) {
	if i.Field.MilestoneID == nil {
		return nil, nil
	}

	milestone, err := i.Resolver.MilestoneService.FindMilestoneByID(*i.Field.MilestoneID)
	if err != nil {
		return nil, err
	}

	return &MilestoneResolver{
		Field:    milestone,
		Resolver: i.Resolver,
	}, nil
}

// Author resolve ...
func (i *IssueResolver) Author() (*UserResolver, error) {
	user, err := i.Resolver.UserService.FindUserByID(i.Field.AuthorID)
	if err != nil {
		return nil, err
	}

	return &UserResolver{
		Field:    user,
		Resolver: i.Resolver,
	}, nil
}

// Number resolve ...
func (i *IssueResolver) Number() int32 {
	return helpers.IntToInt32(i.Field.Number)
}

// Title resolve ...
func (i *IssueResolver) Title() string {
	return i.Field.Title
}

// Body resolve ...
func (i *IssueResolver) Body() *string {
	return i.Field.Body
}

// State is to resolve ...
func (i *IssueResolver) State() string {
	return i.Field.State
}

// IssueType is resolve ...
func (i *IssueResolver) IssueType() string {
	return i.Field.IssueType
}

// IssueConnectionResolver is to resolve
type IssueConnectionResolver struct {
	Issues    []*IssueResolver
	Paginator repository.PageInfo
	Count     int
}

// Edges resolve ...
func (i *IssueConnectionResolver) Edges() *[]*IssueEdgeResolver {
	issues := make([]*IssueEdgeResolver, len(i.Issues))

	for i, issue := range i.Issues {
		issues[i] = &IssueEdgeResolver{
			cursor: issue.ID(),
			Field:  issue,
		}
	}

	return &issues
}

// Nodes to get available nodes
func (i *IssueConnectionResolver) Nodes() *[]*IssueResolver {
	return &i.Issues
}

// TotalCount to get project total
func (i *IssueConnectionResolver) TotalCount() int32 {
	return int32(i.Count)
}

// PageInfo is resolves ...
func (i *IssueConnectionResolver) PageInfo() PageInfoResolver {
	return PageInfoResolver{
		startCursor: helpers.String(helpers.EncryptID("issue", i.Paginator.StartCursor)),
		endCursor:   helpers.String(helpers.EncryptID("issue", i.Paginator.EndCursor)),
	}
}

// IssueEdgeResolver is resolve ...
type IssueEdgeResolver struct {
	cursor graphql.ID
	Field  *IssueResolver
}

// Cursor resolve ...
func (i *IssueEdgeResolver) Cursor() graphql.ID {
	return graphql.ID(i.Field.ID())
}

// Node to resolve
func (i *IssueEdgeResolver) Node() *IssueResolver {
	return i.Field
}
