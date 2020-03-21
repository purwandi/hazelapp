package resolver

import (
	"context"

	"github.com/purwandi/hazelapp/project/types"
)

type GetProjectsInput struct {
	First  *string
	Last   *string
	After  string
	Before string
}

func (r *Resolver) Projects(ctx context.Context, args struct{ Input GetProjectsInput }) (*ProjectConnectionResolver, error) {
	input := types.GetProjectsInput{}

	result, err := r.ProjectService.GetProjects(input)
	if err != nil {
		return nil, err
	}

	projects := make([]*ProjectResolver, len(result))
	for i, project := range result {
		projects[i] = &ProjectResolver{
			Field:    project,
			Resolver: r,
		}
	}

	return &ProjectConnectionResolver{
		Projects: projects,
		Count:    len(result),
	}, nil
}

// func (r *Resolver) Project(ctx context.Context, args struct{ ID graphql.ID }) (*ProjectResolver, error) {
// 	uid, err := uuid.FromString(string(args.ID))
// 	if err != nil {
// 		return nil, err
// 	}

// 	result, err := r.ProjectService.FindProjectByID(uid)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &ProjectResolver{
// 		Field:    result,
// 		Resolver: r,
// 	}, nil
// }
