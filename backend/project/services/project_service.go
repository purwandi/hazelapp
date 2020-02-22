package services

import (
	"github.com/purwandi/hazelapp/project/domain"
	"github.com/purwandi/hazelapp/project/repository"
)

type ProjectService struct {
	query      repository.Query
	repository repository.Repository
}

func NewProjectService(query repository.Query, repo repository.Repository) *ProjectService {
	return &ProjectService{
		query:      query,
		repository: repo,
	}
}

func (s *ProjectService) CreateProject(name, description string) (domain.Project, error) {
	// process
	project, err := domain.CreateProject(name, description)
	if err != nil {
		return domain.Project{}, err
	}

	// persist
	err = <-s.repository.Save(project)
	if err != nil {
		return domain.Project{}, err
	}

	// response
	return domain.Project{}, nil
}

func (s *ProjectService) FindAllProject() ([]domain.Project, error) {
	// process
	result := <-s.query.FindAll()
	if result.Error != nil {
		return []domain.Project{}, result.Error
	}

	// response
	return result.Result.([]domain.Project), nil
}
