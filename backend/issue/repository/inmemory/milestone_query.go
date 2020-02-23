package inmemory

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/storage"
	uuid "github.com/satori/go.uuid"
)

type MilestoneQueryInMemory struct {
	Storage *storage.MilestoneStorage
}

func NewMilestoneQueryInMemory(s *storage.MilestoneStorage) repository.MilestoneQuery {
	return &MilestoneQueryInMemory{Storage: s}
}

func (query *MilestoneQueryInMemory) FindAllMilestonesByProjectID(id uuid.UUID) <-chan repository.QueryResult {
	result := make(chan repository.QueryResult)

	go func() {
		milestones := []domain.Milestone{}
		for _, milestone := range query.Storage.MilestoneMap {
			if milestone.ProjectID == id {
				milestones = append(milestones, milestone)
			}
		}

		result <- repository.QueryResult{Result: milestones}
		close(result)
	}()

	return result
}
