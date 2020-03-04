package resolver

import (
	"context"

	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/project/types"
	"github.com/sirupsen/logrus"
)

// CreateProjectInput is an graphql input contract to create new project
type CreateProjectInput struct {
	OwnerID     *int32 `json:"ownerId"`
	Name        string
	Description *string
}

// ProjectCreate to create new project
func (r *Resolver) ProjectCreate(ctx context.Context, args struct{ Input CreateProjectInput }) (*ProjectResolver, error) {
	// TODO : Neet to change to current user id if not present
	input := types.CreateProjectInput{
		OwnerID:     int(helpers.Int32Value(args.Input.OwnerID)),
		Name:        args.Input.Name,
		Description: args.Input.Description,
	}

	project, err := r.ProjectService.CreateProject(input)
	logrus.Info(project)

	if err != nil {
		return nil, err
	}

	return &ProjectResolver{
		Field:    project,
		Resolver: r,
	}, nil
}
