package types

// CreateIssueInput is to create new issue input
type CreateIssueInput struct {
	ProjectID   int
	MilestoneID int
	AuthorID    int
	Title       string
	Body        string
}

// IssueOrder ...
type IssueOrder struct {
	Field     string
	Direction string
}

// IssueFilter ...
type IssueFilter struct {
	Field string
	Value interface{}
}

// GetIssuesInput ...
type GetIssuesInput struct {
	ProjectID    int
	IssueFilters *[]IssueFilter
	First        *int
	Last         *int
	After        *int
	Before       *int
	IssueOrder   *IssueOrder
}

// GetIssueInput input ...
type GetIssueInput struct {
	ProjectID int
	Number    int
}
