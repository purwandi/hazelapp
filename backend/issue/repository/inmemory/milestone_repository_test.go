package inmemory

import (
	"testing"

	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/storage"
	"github.com/stretchr/testify/assert"
)

func TestCanSaveMilestone(t *testing.T) {
	// Given
	milestone := domain.Milestone{
		ProjectID:   1,
		Name:        "1.2.0",
		Description: "Goal to create best milestone",
	}

	store := storage.NewMilestoneStorage()
	repo := NewMilestoneRepositoryInMemory(store)

	// When
	result := <-repo.Save(&milestone)

	// Then
	assert.NoError(t, result)
	assert.Equal(t, 1, milestone.ID)
}

func TestCanIncrementMilestoneID(t *testing.T) {
	// Given
	milestone1 := domain.Milestone{
		ProjectID:   1,
		Name:        "1.2.0",
		Description: "Goal to create best milestone",
	}
	milestone2 := domain.Milestone{
		ProjectID:   1,
		Name:        "1.2.0",
		Description: "Goal to create best milestone",
	}

	store := storage.NewMilestoneStorage()
	repo := NewMilestoneRepositoryInMemory(store)

	// When
	<-repo.Save(&milestone1)
	result := <-repo.Save(&milestone2)

	// Then
	assert.NoError(t, result)
	assert.Equal(t, 1, milestone1.ID)
	assert.Equal(t, 2, milestone2.ID)
}
