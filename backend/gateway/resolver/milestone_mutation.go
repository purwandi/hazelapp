package resolver

import (
	"context"

	"github.com/purwandi/hazelapp/types"
	uuid "github.com/satori/go.uuid"
)

type CreateMilestoneInput struct {
	Name        string
	Description *string
	ProjectId   string
}

func (r *Resolver) MilestoneCreate(ctx context.Context, args struct{ Input CreateMilestoneInput }) (*MilestoneResolver, error) {
	projectID, err := uuid.FromString(args.Input.ProjectId)
	if err != nil {
		return nil, err
	}

	// Process
	milestone, err := r.MilestoneService.CreateMilestone(
		args.Input.Name,
		types.StringValue(args.Input.Description),
		projectID,
	)

	if err != nil {
		return nil, err
	}

	return &MilestoneResolver{
		Field:    milestone,
		Resolver: r,
	}, nil
}
