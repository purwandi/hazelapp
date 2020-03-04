package domain

type ProjectError struct {
	Code int
}

const (
	ProjectErrorNameIsBlank = iota
	ProjectErrorNameIsExists
)

func (e ProjectError) Error() string {
	switch e.Code {
	default:
		return "Unrecognized project error code"
	case ProjectErrorNameIsBlank:
		return "Project name can't be blank"
	case ProjectErrorNameIsExists:
		return "Project name already exists"
	}
}
