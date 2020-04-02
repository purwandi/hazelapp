package storage

import "github.com/purwandi/hazelapp/issue/domain"

// MilestoneStorage is milestone data storage in memory
type MilestoneStorage struct {
	MilestoneMap []domain.Milestone
}

// NewMilestoneStorage is to create new MilestoneStorage instance
func NewMilestoneStorage() *MilestoneStorage {
	return &MilestoneStorage{
		MilestoneMap: []domain.Milestone{},
	}
}

// Demo is for demo
func (m *MilestoneStorage) Demo() error {
	m.MilestoneMap = []domain.Milestone{
		domain.Milestone{ID: 1, ProjectID: 3, Name: "Sprint 1"},
		domain.Milestone{ID: 2, ProjectID: 2, Name: "Sprint February"},
	}
	return nil
}
