package inmemory

import (
	"testing"

	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/storage"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestCanGetLastIssueNumberFromGivenProject(t *testing.T) {
	// Given
	issueID, _ := uuid.NewV4()
	projectID, _ := uuid.NewV4()

	store := storage.NewIssueStorage()
	store.IssueMap = []domain.Issue{
		domain.Issue{ID: issueID, Title: "Hello", Body: "This is a body", AutoNumber: 1, ProjectID: projectID},
	}

	query := NewIssueQueryInMemory(store)

	// When
	result := <-query.GetLastIssueNumberFromGivenProject(projectID)

	// Then
	assert.NoError(t, result.Error)
	assert.Equal(t, 1, result.Result.(int))
}
