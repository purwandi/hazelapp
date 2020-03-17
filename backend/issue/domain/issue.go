package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

const (
	Open   = "OPEN"
	Closed = "CLOSED"
)

type Issue struct {
	ID        uuid.UUID
	ProjectID uuid.UUID
	AuthorID  uuid.UUID
	IID       int
	Title     string
	Body      string
	State     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type IssueService interface {
	GetLastIssueNumberFromGivenProject(uuid.UUID) (int, error)
}

func CreateIssue(service IssueService, title, body string, projectID uuid.UUID) (*Issue, error) {
	if title == "" {
		return nil, IssueError{IssueErrorTitleIsBlank}
	}

	iid, err := service.GetLastIssueNumberFromGivenProject(projectID)
	if err != nil {
		return nil, err
	}

	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	return &Issue{
		ID:        uid,
		ProjectID: projectID,
		IID:       iid + 1,
		Title:     title,
		Body:      body,
		State:     Open,
		CreatedAt: time.Now(),
	}, nil
}

func (i *Issue) ChangeTitle(title string) error {
	if title == "" {
		return IssueError{IssueErrorTitleIsBlank}
	}

	i.Title = title
	i.UpdatedAt = time.Now()
	return nil
}

func (i *Issue) ChangeBody(body string) error {
	i.Body = body
	i.UpdatedAt = time.Now()
	return nil
}

func (i *Issue) ChangeState(state string) error {
	if state != Open {
		return IssueError{IssueErrorUndefinedIssueState}
	}

	if state != Closed {
		return IssueError{IssueErrorUndefinedIssueState}
	}

	i.State = state
	i.UpdatedAt = time.Now()
	return nil
}

func (i *Issue) ReOpen() error {
	i.State = Open
	i.UpdatedAt = time.Now()
	return nil
}

func (i *Issue) Close() error {
	i.State = Closed
	i.UpdatedAt = time.Now()
	return nil
}
