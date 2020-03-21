package resolver

import (
	"context"

	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/project/types"
	"github.com/purwandi/hazelapp/user/domain"
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
	user := ctx.Value("viewer").(domain.User)
	input := types.CreateProjectInput{
		Name:        args.Input.Name,
		Description: args.Input.Description,
	}

	if args.Input.OwnerID != nil {
		input.OwnerID = int(helpers.Int32Value(args.Input.OwnerID))
	} else {
		input.OwnerID = user.ID
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
