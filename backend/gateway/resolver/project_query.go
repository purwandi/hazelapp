package resolver

import (
	"context"

	"github.com/graph-gophers/graphql-go"
	uuid "github.com/satori/go.uuid"
)

func (r *Resolver) Projects(ctx context.Context) (*ProjectConnectionResolver, error) {
	result, err := r.ProjectService.FindAllProject()
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

func (r *Resolver) Project(ctx context.Context, args struct{ ID graphql.ID }) (*ProjectResolver, error) {
	uid, err := uuid.FromString(string(args.ID))
	if err != nil {
		return nil, err
	}

	result, err := r.ProjectService.FindProjectByID(uid)
	if err != nil {
		return nil, err
	}

	return &ProjectResolver{
		Field:    result,
		Resolver: r,
	}, nil
}
