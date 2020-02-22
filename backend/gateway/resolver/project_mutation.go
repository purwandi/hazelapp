package resolver

import (
	"context"

	"github.com/purwandi/hazelapp/gateway/types"
	"github.com/sirupsen/logrus"
)

type ProjectCreateInput struct {
	Name        string
	Description *string
}

func (r *Resolver) ProjectCreate(ctx context.Context, args ProjectCreateInput) (*ProjectResolver, error) {
	project, err := r.ProjectService.CreateProject(args.Name, types.StringValue(args.Description))
	logrus.Info(project)

	if err != nil {
		return nil, err
	}

	return &ProjectResolver{
		Field:    project,
		Resolver: r,
	}, nil
}
