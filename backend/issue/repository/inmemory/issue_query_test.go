package inmemory

import (
	"testing"

	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/storage"
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
