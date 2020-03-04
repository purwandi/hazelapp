package inmemory

import (
	"testing"

	"github.com/purwandi/hazelapp/helpers"
	"github.com/purwandi/hazelapp/user/domain"
	"github.com/purwandi/hazelapp/user/storage"
	"github.com/purwandi/hazelapp/user/types"
	"github.com/stretchr/testify/assert"
)

func TestCanFindUser(t *testing.T) {
	// Given
	store := storage.NewUserStorage()
	store.Demo()

	query := NewUserQueryInMemory(store)

	// When find user by id
	result := <-query.FindUser(&types.FindUserInput{ID: helpers.Int(1)})
	assert.Equal(t, 1, result.Result.(domain.User).ID)

	// When find user by email
	result = <-query.FindUser(&types.FindUserInput{Email: helpers.String("foobar@kaboor.com")})
	assert.Equal(t, 1, result.Result.(domain.User).ID)

	// When find user by username
	result = <-query.FindUser(&types.FindUserInput{Username: helpers.String("foobar")})
	assert.Equal(t, 1, result.Result.(domain.User).ID)
}
