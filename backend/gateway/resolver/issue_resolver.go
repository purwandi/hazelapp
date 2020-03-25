package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/issue/domain"
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
	if i.Field.MilestoneID == 0 {
		return nil, nil
	}

	milestone, err := i.Resolver.MilestoneService.FindMilestoneByID(i.Field.MilestoneID)
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
	return helpers.String(i.Field.Body)
}
