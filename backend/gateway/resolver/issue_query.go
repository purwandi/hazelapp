package resolver

import (
	"context"

	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/issue/types"
)

// GetIssuesInput is get issues input
type GetIssuesInput struct {
	First  *int32
	Last   *int32
	After  *string
	Before *string
}

// Issues is query
func (r *Resolver) Issues(ctx context.Context, projectID int, args struct{ Input GetIssuesInput }) (*IssueConnectionResolver, error) {
	var after *int
	var before *int

	// TODO refactor into separated function
	if args.Input.After != nil {
		inputAfter, err := helpers.DecryptID("issue", *args.Input.After)
		if err != nil {
			return nil, err
		}
		after = &inputAfter
	}

	// TODO refactor into separated function
	if args.Input.Before != nil {
		inputBefore, err := helpers.DecryptID("issue", *args.Input.Before)
		if err != nil {
			return nil, err
		}
		before = &inputBefore
	}

	input := types.GetIssuesInput{
		ProjectID: projectID,
		First:     helpers.Int32ValueToInt(args.Input.First),
		Last:      helpers.Int32ValueToInt(args.Input.Last),
		After:     after,
		Before:    before,
	}

	result, err := r.IssueService.GetIssues(input)
	if err != nil {
		return &IssueConnectionResolver{}, err
	}

	issues := make([]*IssueResolver, len(result.Issues))
	for i, item := range result.Issues {
		issues[i] = &IssueResolver{
			Field:    item,
			Resolver: r,
		}
	}

	return &IssueConnectionResolver{
		Count:     result.Total,
		Issues:    issues,
		Paginator: result.PageInfo,
	}, nil
}
