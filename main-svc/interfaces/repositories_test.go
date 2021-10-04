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

func createUser(t *testing.T) (*domain.User, error) {
	t.Helper()
	gdb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	assert.NoError(t, err)

	userRepo := interfaces.UserRepo{
		DB: gdb,
	}

	u := domain.User{
		Name:     "John Doe",
		Username: "johndoe",
		Password: "johndoe123",
	}

	user, err := userRepo.Store(u)
	assert.NoError(t, err)

	t.Cleanup(func() {
		err := userRepo.Delete(user.ID)
		assert.NoError(t, err)
	})

	return user, nil
}

func TestUser_FindByID(t *testing.T) {
	gdb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	assert.NoError(t, err)

	userRepo := interfaces.UserRepo{
		DB: gdb,
	}
	user, err := createUser(t)
	assert.NoError(t, err)

	userGot, err := userRepo.FindByID(user.ID)
	assert.NoError(t, err)

	want := &domain.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}

	got := userGot

	assert.Equal(t, want, got)
}

func TestStoreUser(t *testing.T) {
	gdb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	assert.NoError(t, err)

	userRepo := interfaces.UserRepo{
		DB: gdb,
	}

	u, err := createUser(t)
	assert.NoError(t, err)

	user, err := userRepo.FindByID(u.ID)
	assert.NoError(t, err)
	assert.Equal(t, "johndoe", user.Username)
}

func TestUser_FindByUsername(t *testing.T) {
	gdb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	assert.NoError(t, err)

	userRepo := interfaces.UserRepo{
		DB: gdb,
	}

	u, err := createUser(t)
	assert.NoError(t, err)

	user, err := userRepo.FindByUsername(u.Username)
	assert.NoError(t, err)

	want := &domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Username: u.Username,
	}

	got := user

	assert.Equal(t, want, got)
}
