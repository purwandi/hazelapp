package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanEncryptID(t *testing.T) {
	// Given
	id := 1

	// When
	result := EncryptID("user", id)

	// Then
	assert.Equal(t, "dXNlci0x", result)
}

func TestCanDecrypt(t *testing.T) {
	// Given
	id := 1
	r := EncryptID("user", id)

	// When
	result, _ := DecryptID("user", r)

	// Then
	assert.Equal(t, 1, result)
}
