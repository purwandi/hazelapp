package types

import "time"

// GetMilestonesInput is input to find milestone from given project id
type GetMilestonesInput struct {
	ProjectID int
}

// CreateMilestoneInput is input to create milestone
type CreateMilestoneInput struct {
	ProjectID   int
	Name        string
	Description *string
	StartedAt   *time.Time
	EndedAt     *time.Time
}
