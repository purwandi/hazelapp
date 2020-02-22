package storage

import (
	"time"

	"github.com/purwandi/hazelapp/project/domain"
	uuid "github.com/satori/go.uuid"
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
	uid1, _ := uuid.FromString("499acd21-a683-49c8-a404-dc04fb8d3112")
	uid2, _ := uuid.FromString("3a9a774b-fa6d-4440-8b94-db5f37cd32a4")
	uid3, _ := uuid.FromString("c4d33471-65b5-4406-9e13-e277ada92575")
	uid4, _ := uuid.FromString("aaa58e7f-b2d1-415a-b406-d0ca80b32b55")
	storage.ProjectMap = []domain.Project{
		domain.Project{ID: uid1, Name: "Project Website", Description: "Add new features", CreatedAt: time.Now()},
		domain.Project{ID: uid2, Name: "Marketing Website", Description: "Implement project management for linear issues", CreatedAt: time.Now()},
		domain.Project{ID: uid3, Name: "Mobile App", Description: "Robus filtering on all views", CreatedAt: time.Now()},
		domain.Project{ID: uid4, Name: "Cycles", Description: "Initial version of cycles", CreatedAt: time.Now()},
	}

	return nil
}
