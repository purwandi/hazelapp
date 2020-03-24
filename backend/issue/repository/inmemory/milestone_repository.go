package inmemory

import (
	"github.com/purwandi/hazelapp/issue/domain"
	"github.com/purwandi/hazelapp/issue/repository"
	"github.com/purwandi/hazelapp/issue/storage"
)

// MilestoneRepositoryInMemory is milestone reopsitory implementation in memory
type MilestoneRepositoryInMemory struct {
	Storage *storage.MilestoneStorage
}

// NewMilestoneRepositoryInMemory is to create MilestoneRepositoryInMemory instance
func NewMilestoneRepositoryInMemory(s *storage.MilestoneStorage) repository.MilestoneRepository {
	return &MilestoneRepositoryInMemory{Storage: s}
}

// Save for save
func (repo *MilestoneRepositoryInMemory) Save(milestone *domain.Milestone) <-chan error {
	result := make(chan error)

	go func() {
		milestone.ID = 1
		if len(repo.Storage.MilestoneMap) > 0 {
			lastID := repo.Storage.MilestoneMap[len(repo.Storage.MilestoneMap)-1].ID
			milestone.ID = lastID + 1
		}

		repo.Storage.MilestoneMap = append(repo.Storage.MilestoneMap, *milestone)

		result <- nil
		close(result)
	}()

	return result
}
