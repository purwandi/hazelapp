package storage

import "github.com/purwandi/hazelapp/issue/domain"

type IssueStorage struct {
	IssueMap []domain.Issue
}

func NewIssueStorage() *IssueStorage {
	return &IssueStorage{
		IssueMap: []domain.Issue{},
	}
}
