package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Project struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func CreateProject(name, description string) (*Project, error) {
	if name == "" {
		return nil, ProjectError{ProjectErrorNameIsBlank}
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	return &Project{
		ID:          uid,
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
	}, nil
}
