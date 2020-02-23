package domain

type MilestoneError struct {
	Code int
}

const (
	MilestoneErrorNameIsBlank = iota
)

func (e MilestoneError) Error() string {
	switch e.Code {
	default:
		return "Unrecognized milestone error code"
	case MilestoneErrorNameIsBlank:
		return "Milestone name can'be blank"
	}
}
