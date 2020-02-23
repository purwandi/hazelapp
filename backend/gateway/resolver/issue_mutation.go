package resolver

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

type IssueCreateInput struct {
	Title     string
	Body      string
	ProjectID string
}

func (r *Resolver) IssueCreate(ctx context.Context, args IssueCreateInput) (*IssueResolver, error) {
	projectID, err := uuid.FromString(args.ProjectID)
	if err != nil {
		return nil, err
	}

	issue, err := r.IssueService.CreateIssue(args.Title, args.Body, projectID)
	if err != nil {
		return nil, err
	}

	return &IssueResolver{
		Field:    issue,
		Resolver: r,
	}, nil
}
