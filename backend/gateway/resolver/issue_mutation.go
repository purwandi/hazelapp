package resolver

import (
	"context"

	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/issue/types"
	"github.com/purwandi/hazelapp/user/domain"
)

// IssueCreateInput is to wrap issue input
type IssueCreateInput struct {
	ProjectID   string
	MilestoneID *string
	Title       string
	Body        *string
}

// IssueCreate to create an isue
func (r *Resolver) IssueCreate(ctx context.Context, args struct{ Input IssueCreateInput }) (*IssueResolver, error) {
	viewer := ctx.Value("viewer").(domain.User)

	projectID, err := helpers.DecryptID("project", args.Input.ProjectID)
	if err != nil {
		return nil, err
	}

	project, err := r.ProjectService.FindProjectByID(projectID)
	if err != nil {
		return nil, err
	}

	input := types.CreateIssueInput{
		ProjectID: project.ID,
		AuthorID:  viewer.ID,
		Title:     args.Input.Title,
	}

	// Set milestone if milestone is set
	if args.Input.MilestoneID != nil {
		milestoneID, err := helpers.DecryptID("milestone", *args.Input.MilestoneID)
		if err != nil {
			return nil, err
		}

		milestone, err := r.MilestoneService.FindMilestoneByID(milestoneID)
		if err != nil {
			return nil, err
		}

		input.MilestoneID = milestone.ID
	}

	if args.Input.Body != nil {
		input.Body = *args.Input.Body
	}

	issue, err := r.IssueService.CreateIssue(input)
	if err != nil {
		return nil, err
	}

	return &IssueResolver{
		Field:    issue,
		Resolver: r,
	}, nil
}
