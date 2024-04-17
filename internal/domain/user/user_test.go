package user_test

import (
	"testing"

	"github.com/josue/chalenge_golang/internal/domain/user"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	userName := "testUser"
	email := "test@example.com"
	phone := "1234567890"
	password := "password123"

	user := user.NewUser(userName, email, phone, password)
	assert.Equal(t, userName, user.UserName)
	assert.Equal(t, email, user.Email)
	assert.Equal(t, phone, user.PhoneNumber)
	assert.Equal(t, password, user.Pass)

}
