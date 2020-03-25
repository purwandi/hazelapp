package domain

// IssueError is wrap issue error message
type IssueError struct {
	Code int
}

const (
	// IssueErrorTitleIsBlank when title is blank
	IssueErrorTitleIsBlank = iota
	// IssueErrorUndefinedIssueState when issue state is not correct
	IssueErrorUndefinedIssueState
)

func (e IssueError) Error() string {
	switch e.Code {
	default:
		return "Unreconized issue error code"
	case IssueErrorTitleIsBlank:
		return "Title can't be blank"
	case IssueErrorUndefinedIssueState:
		return "Undefined issue state, state only accept `CLOSED` or `OPEN`"
	}
}
