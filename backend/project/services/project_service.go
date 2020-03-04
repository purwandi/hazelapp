package services

import (
	"github.com/purwandi/hazelapp/project/domain"
	"github.com/purwandi/hazelapp/project/repository"
	"github.com/purwandi/hazelapp/project/types"
	"github.com/sirupsen/logrus"
)

// ProjectService is to wrap project service domain
type ProjectService struct {
	query      repository.ProjectQuery
	repository repository.ProjectRepository
}

// NewProjectService is to create new ProjectService instance
func NewProjectService(query repository.ProjectQuery, repo repository.ProjectRepository) *ProjectService {
	return &ProjectService{
		query:      query,
		repository: repo,
	}
}

// FindProject is to find project
func (s *ProjectService) FindProject(args types.FindProjectInput) (domain.Project, error) {
	result := <-s.query.FindProject(args)
	if result.Error != nil {
		return domain.Project{}, result.Error
	}

	project := result.Result.(domain.Project)
	if project.ID == 0 {
		return domain.Project{}, ServiceError{ServiceErrorResourceNotFound}
	}

	return result.Result.(domain.Project), nil
}

// CreateProject is to create new project
func (s *ProjectService) CreateProject(args types.CreateProjectInput) (domain.Project, error) {
	// process
	project, err := domain.CreateProject(s, args)
	if err != nil {
		return domain.Project{}, err
	}

	// persist
	err = <-s.repository.Save(project)
	if err != nil {
		return domain.Project{}, err
	}

	logrus.Info(project)

	// response
	return *project, nil
}

// func (s *ProjectService) FindAllProject() ([]domain.Project, error) {
// 	// process
// 	result := <-s.query.FindAll()
// 	if result.Error != nil {
// 		return []domain.Project{}, result.Error
// 	}

// 	// response
// 	return result.Result.([]domain.Project), nil
// }

// func (s *ProjectService) FindProjectByID(id uuid.UUID) (domain.Project, error) {
// 	// process
// 	result := <-s.query.FindProjectByID(id)
// 	if result.Error != nil {
// 		return domain.Project{}, result.Error
// 	}

// 	// response
// 	return result.Result.(domain.Project), nil
// }
