package domain

import (
	"testing"

	"github.com/purwandi/hazelapp/issue/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type IssueServiceMock struct {
	mock.Mock
}

func (mock IssueServiceMock) GetLastIssueNumberFromGivenProject(projectID int) (int, error) {
	args := mock.Called(projectID)
	return args.Int(0), args.Error(1)
}

func TestCanCreateIssue(t *testing.T) {
	// Given
	input := types.CreateIssueInput{
		ProjectID: 1,
		AuthorID:  1,
		Title:     "New project with torth",
		IssueType: "story",
	}

	service := IssueServiceMock{}
	service.On("GetLastIssueNumberFromGivenProject", 1).Return(2, nil)

	// When
	issue, error := CreateIssue(service, input)

	// Then
	assert.NoError(t, error)
	assert.Equal(t, 1, issue.ProjectID)
	assert.Nil(t, issue.MilestoneID)
	assert.Equal(t, 3, issue.Number)
	assert.Equal(t, input.Title, issue.Title)
	assert.Equal(t, IssueOpen, issue.State)
}

func TestCanAssignMilestone(t *testing.T) {
	// Given
	issue := Issue{
		ProjectID: 1,
		AuthorID:  1,
		Title:     "Hello world",
	}

	//
}

func TestCanChangeIssueTitle(t *testing.T) {
	// Given
	issue := &Issue{
		Title: "New issue from check",
	}

	// When
	issue.ChangeTitle("hallllo")

	// Then
	assert.Equal(t, "hallllo", issue.Title)
}

func TestCanUnableToChangeIssueTitle(t *testing.T) {
	// Given
	issue := &Issue{
		Title: "New issue from check",
	}

	// When
	issue.ChangeTitle("")

	// Then
	assert.Equal(t, "New issue from check", issue.Title)
}

func TestCanUnableToChangeIssueState(t *testing.T) {
	// Given
	issue := &Issue{
		Title: "New issue from check",
		State: IssueOpen,
	}

	// When
	issue.Closed()

	// Then
	assert.Equal(t, IssueClosed, issue.State)

	// When
	issue.ReOpen()

	// Then
	assert.Equal(t, IssueOpen, issue.State)
}
