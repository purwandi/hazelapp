package types

// CreateIssueInput is to create new issue input
type CreateIssueInput struct {
	ProjectID   int
	MilestoneID int
	AuthorID    int
	Title       string
	Body        string
}

// GetIssuesInput to issues
type GetIssuesInput struct {
	ProjectID int
}

type GetIssueInput struct {
	ProjectID int
	Number    int
}
