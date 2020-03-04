package inmemory

import (
	"testing"

	"github.com/purwandi/hazelapp/user/domain"
	"github.com/purwandi/hazelapp/user/storage"
	"github.com/stretchr/testify/assert"
)

func TestCanSaveUser(t *testing.T) {
	// Given
	user := domain.User{Email: "foobar@ba.com"}
	store := storage.NewUserStorage()
	repo := NewUserRepositoryInMemory(store)

	// When
	result := <-repo.Save(&user)

	// Then
	assert.Equal(t, 1, user.ID)
	assert.NoError(t, result)
}

func TestCanIncrementUserID(t *testing.T) {
	// Given
	user := domain.User{Email: "foobar@ba.com"}
	store := storage.NewUserStorage()
	store.UserMap = []domain.User{
		domain.User{ID: 1, Email: "foobar@ko.com"},
	}

	repo := NewUserRepositoryInMemory(store)

	// When
	result := <-repo.Save(&user)

	// Then
	assert.Equal(t, 2, user.ID)
	assert.NoError(t, result)
}
