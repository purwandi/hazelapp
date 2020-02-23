package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Issue struct {
	ID         uuid.UUID
	ProjectID  uuid.UUID
	AutoNumber int
	Title      string
	Body       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type IssueService interface {
	GetLastIssueNumberFromGivenProject(uuid.UUID) (int, error)
}

func CreateIssue(service IssueService, title, body string, projectID uuid.UUID) (*Issue, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	autonumber, err := service.GetLastIssueNumberFromGivenProject(projectID)
	if err != nil {
		return nil, err
	}

	return &Issue{
		ID:         uid,
		ProjectID:  projectID,
		AutoNumber: autonumber + 1,
		Title:      title,
		Body:       body,
		CreatedAt:  time.Now(),
	}, nil
}
