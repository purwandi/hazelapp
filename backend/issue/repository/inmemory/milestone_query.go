package inmemory

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/storage"
	"github.com/purwandi/hazelapp/issue/types"
)

// MilestoneQueryInMemory is MilestoneQuery implementation in memory
type MilestoneQueryInMemory struct {
	Storage *storage.MilestoneStorage
}

// NewMilestoneQueryInMemory to create new MilestoneQueryInMemory instance
func NewMilestoneQueryInMemory(s *storage.MilestoneStorage) repository.MilestoneQuery {
	return &MilestoneQueryInMemory{Storage: s}
}

// GetMilestones to get milestones
func (query *MilestoneQueryInMemory) GetMilestones(args types.GetMilestonesInput) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		milestones := []domain.Milestone{}
		for _, milestone := range query.Storage.MilestoneMap {
			if milestone.ProjectID == args.ProjectID {
				milestones = append(milestones, milestone)
			}
		}

		result <- repository.QueryResult{Result: milestones}
		close(result)
	}()

	return result
}

// FindMilestoneByID to get milestones
func (query *MilestoneQueryInMemory) FindMilestoneByID(id int) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		milestone := domain.Milestone{}
		for _, item := range query.Storage.MilestoneMap {
			if item.ID == id {
				milestone = item
			}
		}

		result <- repository.QueryResult{Result: milestone}
		close(result)
	}()

	return result
}
