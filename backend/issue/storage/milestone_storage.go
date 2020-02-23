package storage

import "github.com/purwandi/hazelapp/issue/domain"

type MilestoneStorage struct {
	MilestoneMap []domain.Milestone
}

func NewMilestoneStorage() *MilestoneStorage {
	return &MilestoneStorage{
		MilestoneMap: []domain.Milestone{},
	}
}
