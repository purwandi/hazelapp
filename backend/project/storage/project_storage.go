package storage

import (
	"time"

	"github.com/purwandi/hazelapp/project/domain"
)

type ProjectStorage struct {
	ProjectMap []domain.Project
}

func NewProjectStorage() *ProjectStorage {
	return &ProjectStorage{
		ProjectMap: []domain.Project{},
	}
}

func (storage *ProjectStorage) Demo() error {
	storage.ProjectMap = []domain.Project{
		domain.Project{ID: 1, Name: "Project Website", Description: "Add new features", CreatedAt: time.Now()},
		domain.Project{ID: 2, Name: "Marketing Website", Description: "Implement project management for linear issues", CreatedAt: time.Now()},
		domain.Project{ID: 3, Name: "Mobile App", Description: "Robus filtering on all views", CreatedAt: time.Now()},
		domain.Project{ID: 4, Name: "Cycles", Description: "Initial version of cycles", CreatedAt: time.Now()},
	}

	return nil
}
