package services

type ServiceError struct {
	Code int
}

const (
	ServiceErrorResourceNotFound = iota
)

func (e ServiceError) Error() string {
	switch e.Code {
	default:
		return "Unrecognized service error code"
	case ServiceErrorResourceNotFound:
		return "Resource is not found"
	}
}
