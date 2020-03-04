package services

type UserServiceError struct {
	Code int
}

const (
	UserServiceErrorResourceNotFound = iota
)

func (e UserServiceError) Error() string {
	switch e.Code {
	default:
		return "Unrecognized user service error code"
	case UserServiceErrorResourceNotFound:
		return "User resources is not found"
	}
}
