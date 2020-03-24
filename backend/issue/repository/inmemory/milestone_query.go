package inmemory

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/storage"
	"github.com/purwandi/hazelapp/issue/types"
)

type MilestoneQueryInMemory struct {
	Storage *storage.MilestoneStorage
}

func NewMilestoneQueryInMemory(s *storage.MilestoneStorage) repository.MilestoneQuery {
	return &MilestoneQueryInMemory{Storage: s}
}

func (query *MilestoneQueryInMemory) GetMilestones(args *types.GetMilestonesInput) <-chan repository.QueryResult {
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
