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
	// IssueErrorUndefinedIssueType when issue type is not correct
	IssueErrorUndefinedIssueType
)

func (e IssueError) Error() string {
	switch e.Code {
	default:
		return "Unreconized issue error code"
	case IssueErrorTitleIsBlank:
		return "Title can't be blank"
	case IssueErrorUndefinedIssueState:
		return "Undefined issue state, state only accept `CLOSED` or `OPEN`"
	case IssueErrorUndefinedIssueType:
		return "Undefined issue type, issue type only accept `story`, `epic`, `task`, `bug`"
	}
}
