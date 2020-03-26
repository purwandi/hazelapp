package repository

// IssueQueryError is to wrap issue query error
type IssueQueryError struct {
	Code int
}

const (
	// IssueQueryErrorMissingPaginationBoundaries when pagination boundari is not set
	IssueQueryErrorMissingPaginationBoundaries = iota
	// IssueQueryErrorPagination when first and last is passed
	IssueQueryErrorPagination
)

func (e IssueQueryError) Error() string {
	switch e.Code {
	default:
		return "Unrecognized issue query error code"
	case IssueQueryErrorMissingPaginationBoundaries:
		return "Missing pagination boundaries"
	case IssueQueryErrorPagination:
		return "Passing both `first` and `last` to paginate the `Issue` connection is not supported."
	}
}
