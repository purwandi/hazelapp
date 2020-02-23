package domain

type IssueError struct {
	Code int
}

const (
	IssueErrorTitleIsBlank = iota
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
