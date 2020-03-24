package domain

import (
	"time"

	"github.com/purwandi/hazelapp/issue/types"
)

const (
	// MilestoneOpen when milestone is open
	MilestoneOpen = "OPEN"
	// MilestoneClose when milestone is closed
	MilestoneClose = "CLOSED"
)

// Milestone is to wrap milestone model
type Milestone struct {
	ID          int
	ProjectID   int
	Name        string
	Description string
	Status      string
	StartedAt   time.Time
	EndedAt     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CreateMilestone to create milestone
func CreateMilestone(args types.CreateMilestoneInput) (*Milestone, error) {
	if args.Name == "" {
		return nil, MilestoneError{MilestoneErrorNameIsBlank}
	}

	milestone := Milestone{
		ProjectID: args.ProjectID,
		Name:      args.Name,
		Status:    MilestoneOpen,
		CreatedAt: time.Now(),
	}

	if args.Description != nil {
		milestone.Description = *args.Description
	}

	return &milestone, nil
}

// Open to open milestone
func (m *Milestone) Open() error {
	m.Status = MilestoneOpen
	m.UpdatedAt = time.Now()

	return nil
}

// Close to close milestone
func (m *Milestone) Close() error {
	m.Status = MilestoneClose
	m.UpdatedAt = time.Now()

	return nil
}
