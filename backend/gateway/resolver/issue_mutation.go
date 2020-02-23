package resolver

import (
	"context"

	"github.com/purwandi/hazelapp/types"
	uuid "github.com/satori/go.uuid"
)

type IssueCreateInput struct {
	Title     string
	Body      *string
	ProjectId string
}

type IssueUpdateInput struct {
	IID   int32
	Title *string
	Body  *string
	State *string
}

func (r *Resolver) IssueCreate(ctx context.Context, args struct{ Input IssueCreateInput }) (*IssueResolver, error) {
	projectID, err := uuid.FromString(args.Input.ProjectId)
	if err != nil {
		return nil, err
	}

	issue, err := r.IssueService.CreateIssue(
		args.Input.Title,
		types.StringValue(args.Input.Body),
		projectID,
	)

	if err != nil {
		return nil, err
	}

	return &IssueResolver{
		Field:    issue,
		Resolver: r,
	}, nil
}

func (r *Resolver) IssueUpdate(ctx context.Context, args struct{ Input *IssueUpdateInput }) (*IssueResolver, error) {
	return &IssueResolver{}, nil
}
