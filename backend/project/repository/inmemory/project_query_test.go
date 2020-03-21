package inmemory

import (
	"testing"

	"github.com/purwandi/hazelapp/project/domain"
	"github.com/purwandi/hazelapp/project/repository"
	"github.com/purwandi/hazelapp/project/storage"
	"github.com/purwandi/hazelapp/project/types"
	"github.com/stretchr/testify/assert"
)

func TestCanGetFirstProjects(t *testing.T) {
	// Given
	store := storage.NewProjectStorage()
	store.Demo()
	query := NewProjectQueryInMemory(store)
	first := 10

	input := types.GetProjectsInput{
		OwnerID: 1,
		First:   &first,
	}

	// When
	result := <-query.GetProjects(input)

	projects := result.Result.([]domain.Project)

	// Then
	assert.NoError(t, result.Error)
	assert.Equal(t, 10, len(projects))
	assert.Equal(t, 1, projects[0].ID)
	assert.Equal(t, 10, projects[9].ID)
	assert.Equal(t, repository.PageInfo{
		StartCursor: 1,
		EndCursor:   10,
	}, result.PageInfo)
}

func TestCanGetAfterFirstProjects(t *testing.T) {
	// Given
	store := storage.NewProjectStorage()
	store.Demo()
	query := NewProjectQueryInMemory(store)
	first := 6
	after := 10

	input := types.GetProjectsInput{
		OwnerID: 1,
		After:   &after,
		First:   &first,
	}

	// When
	result := <-query.GetProjects(input)

	// Then
	assert.NoError(t, result.Error)
	assert.Equal(t, 6, len(result.Result.([]domain.Project)))
	assert.Equal(t, 11, result.Result.([]domain.Project)[0].ID)
	assert.Equal(t, 16, result.Result.([]domain.Project)[5].ID)
	assert.Equal(t, repository.PageInfo{
		StartCursor: 11,
		EndCursor:   16,
	}, result.PageInfo)
}

func TestCanGetBeforeFirstProjects(t *testing.T) {
	// Given
	store := storage.NewProjectStorage()
	store.Demo()
	query := NewProjectQueryInMemory(store)
	first := 10
	before := 15

	input := types.GetProjectsInput{
		OwnerID: 1,
		First:   &first,
		Before:  &before,
	}

	// When
	result := <-query.GetProjects(input)

	// Then
	assert.NoError(t, result.Error)
	assert.Equal(t, 10, len(result.Result.([]domain.Project)))
	assert.Equal(t, 1, result.Result.([]domain.Project)[0].ID)
	assert.Equal(t, 10, result.Result.([]domain.Project)[9].ID)
	assert.Equal(t, repository.PageInfo{
		StartCursor: 1,
		EndCursor:   10,
	}, result.PageInfo)
}

func TestCanGetBeforeLastProjects(t *testing.T) {
	// Given
	store := storage.NewProjectStorage()
	store.Demo()
	query := NewProjectQueryInMemory(store)
	last := 10
	before := 15

	input := types.GetProjectsInput{
		OwnerID: 1,
		Last:    &last,
		Before:  &before,
	}

	// When
	result := <-query.GetProjects(input)

	// Then
	assert.NoError(t, result.Error)
	assert.Equal(t, 10, len(result.Result.([]domain.Project)))
	assert.Equal(t, 5, result.Result.([]domain.Project)[0].ID)
	assert.Equal(t, 14, result.Result.([]domain.Project)[9].ID)
	assert.Equal(t, repository.PageInfo{
		StartCursor: 5,
		EndCursor:   14,
	}, result.PageInfo)
}

func TestCanNotGetProjects(t *testing.T) {
	// Given
	store := storage.NewProjectStorage()
	query := NewProjectQueryInMemory(store)
	first := 10

	input := types.GetProjectsInput{
		OwnerID: 1,
		First:   &first,
	}

	// When
	result := <-query.GetProjects(input)

	projects := result.Result.([]domain.Project)

	// Then
	assert.NoError(t, result.Error)
	assert.Equal(t, 0, len(projects))
	assert.Equal(t, repository.PageInfo{
		StartCursor: 0,
		EndCursor:   0,
	}, result.PageInfo)
}
