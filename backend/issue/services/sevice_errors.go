package services

type ServiceError struct {
	Code int
}

const (
	ServiceErrorArgsIsBlank = iota
)

func (e ServiceError) Error() string {
	switch e.Code {
	default:
		return "Undefined service error code"
	case ServiceErrorArgsIsBlank:
		return "Service argument can't be be blank"
	}
}
