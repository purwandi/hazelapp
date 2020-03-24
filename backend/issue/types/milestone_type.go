package types

// GetMilestonesInput is input to find milestone from given project id
type GetMilestonesInput struct {
	ProjectID int
}

// CreateMilestoneInput is input to create milestone
type CreateMilestoneInput struct {
	ProjectID   int
	Name        string
	Description *string
}

// CreateIssueInput is to create new issue input
type CreateIssueInput struct {
	ProjectID   int
	MilestoneID int
	AuthorID    int
	Title       string
	Body        string
}
