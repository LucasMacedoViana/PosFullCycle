package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("test", "t@t.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, user.Name, "test")
	assert.Equal(t, user.Email, "t@t.com")
}

func TestUser_ComparePassword(t *testing.T) {
	user, err := NewUser("test", "t@t.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ComparePassword("123456"))
	assert.False(t, user.ComparePassword("1234567"))
	assert.NotEqual(t, "123456", user.Password)
}
