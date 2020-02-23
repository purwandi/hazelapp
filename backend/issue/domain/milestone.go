package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Milestone struct {
	ID          uuid.UUID
	ProjectID   uuid.UUID
	Name        string
	Description string
	StartedAt   time.Time
	EndedAt     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CreateMilestone(name, description string, projectID uuid.UUID) (*Milestone, error) {
	if name == "" {
		return nil, MilestoneError{MilestoneErrorNameIsBlank}
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	return &Milestone{
		ID:          uid,
		ProjectID:   projectID,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}, nil
}
