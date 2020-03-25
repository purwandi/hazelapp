package inmemory

import (
	"testing"

	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/storage"
	"github.com/stretchr/testify/assert"
)

func TestCanSaveIssue(t *testing.T) {
	// Given
	issue := domain.Issue{
		ProjectID: 1,
		Title:     "Helloo",
		Body:      "body",
	}
	store := storage.NewIssueStorage()
	repo := NewIssueRepositoryInMemory(store)

	// When
	result := <-repo.Save(&issue)

	// Then
	assert.NoError(t, result)
	assert.Equal(t, 1, len(store.IssueMap))
	assert.Equal(t, 1, issue.ID)
}

func TestCanIncreaseIssueID(t *testing.T) {
	// Given
	issue := domain.Issue{
		ProjectID: 1,
		Title:     "Helloo",
		Body:      "body",
	}
	store := storage.NewIssueStorage()
	store.IssueMap = []domain.Issue{
		domain.Issue{ID: 2, Number: 3, ProjectID: 1},
	}
	repo := NewIssueRepositoryInMemory(store)

	// When
	result := <-repo.Save(&issue)

	// Then
	assert.NoError(t, result)
	assert.Equal(t, 2, len(store.IssueMap))
	assert.Equal(t, 3, issue.ID)
}
