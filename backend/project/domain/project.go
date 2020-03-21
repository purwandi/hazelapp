package domain

import (
	"time"

	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/project/types"
)

// ProjectService is
type ProjectService interface {
	FindProject(types.FindProjectInput) (Project, error)
}

// Project domain
type Project struct {
	ID          int
	OwnerID     int
	Code        string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CreateProject is to create new project domain instance
func CreateProject(service ProjectService, args types.CreateProjectInput) (*Project, error) {
	if args.Name == "" {
		return nil, ProjectError{ProjectErrorNameIsBlank}
	}

	_, err := service.FindProject(types.FindProjectInput{OwnerID: args.OwnerID, Name: args.Name})
	if err == nil {
		return nil, ProjectError{ProjectErrorNameIsExists}
	}

	return &Project{
		OwnerID:     args.OwnerID,
		Name:        args.Name,
		Code:        helpers.StringValue(args.Code),
		Description: helpers.StringValue(args.Description),
		CreatedAt:   time.Now(),
	}, nil
}
