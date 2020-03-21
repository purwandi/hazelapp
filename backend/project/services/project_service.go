package services

import (
	"github.com/purwandi/hazelapp/project/domain"
	"github.com/purwandi/hazelapp/project/repository"
	"github.com/purwandi/hazelapp/project/types"
)

// ProjectCollectionResult is to wrap project collections result
type ProjectCollectionResult struct {
	Projects []domain.Project
	Total    int
	PageInfo repository.PageInfo
}

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

	// response
	return *project, nil
}

// GetProjects is to get project
func (s *ProjectService) GetProjects(args types.GetProjectsInput) (ProjectCollectionResult, error) {
	// process
	result := <-s.query.GetProjects(args)
	if result.Error != nil {
		return ProjectCollectionResult{}, result.Error
	}

	// response
	return ProjectCollectionResult{
		Projects: result.Result.([]domain.Project),
		Total:    result.Total,
		PageInfo: result.PageInfo,
	}, nil
}
