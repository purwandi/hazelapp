package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/purwandi/hazelapp/issue/types"
)

func TestCanCreateMilestone(t *testing.T) {
	// Given
	input := types.CreateMilestoneInput{
		ProjectID: 1,
		Name:      "1.0.0",
	}

	// When
	milestone, err := CreateMilestone(input)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, "1.0.0", milestone.Name)
	assert.Equal(t, 1, milestone.ProjectID)
	assert.Equal(t, "", milestone.Description)
}
