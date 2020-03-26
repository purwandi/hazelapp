package inmemory

import (
	"testing"

	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/storage"
	"github.com/purwandi/hazelapp/issue/types"
	"github.com/stretchr/testify/assert"
)

func TestCanGetLastIssueNumberFromGivenProject(t *testing.T) {
	// Given
	issueID := 1
	projectID := 2

	store := storage.NewIssueStorage()
	store.IssueMap = []domain.Issue{
		domain.Issue{ID: issueID, Title: "Hello", Body: "This is a body", Number: 1, ProjectID: projectID},
	}

	query := NewIssueQueryInMemory(store)

	// When
	result := <-query.GetLastIssueNumberFromGivenProject(projectID)

	// Then
	assert.NoError(t, result.Error)
	assert.Equal(t, 1, result.Result.(int))
}

func TestCanGetFirstIssues(t *testing.T) {
	// Given
	store := storage.NewIssueStorage()
	store.Demo()

	query := NewIssueQueryInMemory(store)

	// When
	result := <-query.GetIssues(types.GetIssuesInput{
		ProjectID: 1,
		First:     helpers.Int(17),
	})

	// Then
	assert.NoError(t, result.Error)
	assert.Equal(t, 17, len(result.Result.([]domain.Issue)))

	// When
	result = <-query.GetIssues(types.GetIssuesInput{
		ProjectID: 1,
		First:     helpers.Int(9),
	})

	// Then
	assert.Equal(t, 9, len(result.Result.([]domain.Issue)))
}
