package resolver

import (
	"context"
	"time"

	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/issue/types"
)

// CreateMilestoneInput is crate milestone input
type CreateMilestoneInput struct {
	Name        string
	Description *string
	ProjectID   string
	StartDate   *string
	EndDate     *string
}

// MilestoneCreate is to create new milestone
func (r *Resolver) MilestoneCreate(ctx context.Context, args struct{ Input CreateMilestoneInput }) (*MilestoneResolver, error) {
	projectID, err := helpers.DecryptID("project", args.Input.ProjectID)
	if err != nil {
		return nil, err
	}

	project, err := r.ProjectService.FindProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	input := types.CreateMilestoneInput{
		Name:        args.Input.Name,
		Description: args.Input.Description,
		ProjectID:   project.ID,
	}

	if args.Input.StartDate != nil {
		startedAt, err := time.Parse("2006-01-02", *args.Input.StartDate)
		if err != nil {
			return nil, err
		}
		input.StartedAt = &startedAt
	}

	if args.Input.EndDate != nil {
		endedAt, err := time.Parse("2006-01-02", *args.Input.EndDate)
		if err != nil {
			return nil, err
		}
		input.EndedAt = &endedAt
	}

	milestone, err := r.MilestoneService.CreateMilestone(input)
	if err != nil {
		return nil, err
	}

	return &MilestoneResolver{
		Field:    milestone,
		Resolver: r,
	}, nil
}
