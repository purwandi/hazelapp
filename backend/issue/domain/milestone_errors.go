package domain

// MilestoneError is to wrap milestone error
type MilestoneError struct {
	Code int
}

const (
	// MilestoneErrorNameIsBlank when milestone name is blank
	MilestoneErrorNameIsBlank = iota
)

// Error is error interface
func (e MilestoneError) Error() string {
	switch e.Code {
	default:
		return "Unrecognized milestone error code"
	case MilestoneErrorNameIsBlank:
		return "Milestone name can'be blank"
	}
}
