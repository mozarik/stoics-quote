// This is not a unit test because we still depend on open connection to the database

package interfaces_test

import (
	"main-svc/domain"
	"main-svc/interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestUser_FindByID(t *testing.T) {
	gdb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	assert.NoError(t, err)

	userRepo := interfaces.UserRepo{
		DB: gdb,
	}

	user, err := userRepo.FindByID(1)
	assert.NoError(t, err)

	want := &domain.User{
		ID:       1,
		Name:     "John Doe",
		Username: "johndoe",
	}

	got := user

	assert.Equal(t, want, got)
}
