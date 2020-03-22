package resolver

import (
	"context"

	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/project/types"
)

// GetProjectsInput is get projects inputs
type GetProjectsInput struct {
	First  *int32
	Last   *int32
	After  *string
	Before *string
}

// GetProjectInput is to get project input
type GetProjectInput struct {
	Name  string
	Owner string
}

// Projects is get projects
func (r *Resolver) Projects(ctx context.Context, owner int, args struct{ Input *GetProjectsInput }) (*ProjectConnectionResolver, error) {
	var after *int
	var before *int

	// TODO refactor into separated function
	if args.Input.After != nil {
		inputAfter, err := helpers.DecryptID("project", *args.Input.After)
		if err != nil {
			return nil, err
		}
		after = &inputAfter
	}

	// TODO refactor into separated function
	if args.Input.Before != nil {
		inputBefore, err := helpers.DecryptID("project", *args.Input.Before)
		if err != nil {
			return nil, err
		}
		before = &inputBefore
	}

	input := types.GetProjectsInput{
		OwnerID: owner,
		First:   helpers.Int32ValueToInt(args.Input.First),
		Last:    helpers.Int32ValueToInt(args.Input.Last),
		After:   after,
		Before:  before,
	}

	result, err := r.ProjectService.GetProjects(input)
	if err != nil {
		return nil, err
	}

	projects := make([]*ProjectResolver, len(result.Projects))
	for i, project := range result.Projects {
		projects[i] = &ProjectResolver{
			Field:    project,
			Resolver: r,
		}
	}

	return &ProjectConnectionResolver{
		Count:     result.Total,
		Projects:  projects,
		Paginator: result.PageInfo,
	}, nil
}

// Project is resolve ...
func (r *Resolver) Project(ctx context.Context, args GetProjectInput) (*ProjectResolver, error) {
	owner, err := r.UserService.FindUserByUsername(args.Owner)
	if err != nil {
		return nil, err
	}

	input := types.FindProjectInput{OwnerID: owner.ID, Name: args.Name}
	project, err := r.ProjectService.FindProject(input)
	if err != nil {
		return nil, err
	}

	return &ProjectResolver{
		Field:    project,
		Resolver: r,
	}, nil
}
