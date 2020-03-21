package repository

type ProjectQueryError struct {
	Code int
}

const (
	ProjectQueryErrorMissingPaginationBoundaries = iota
	ProjectQueryErrorPagination
)

func (e ProjectQueryError) Error() string {
	switch e.Code {
	default:
		return "Unrecognized project query error code"
	case ProjectQueryErrorMissingPaginationBoundaries:
		return "Missing pagination boundaries"
	case ProjectQueryErrorPagination:
		return "Passing both `first` and `last` to paginate the `project` connection is not supported."
	}
}
