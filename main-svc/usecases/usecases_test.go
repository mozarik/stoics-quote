package usecases_test

import (
	"main-svc/domain"
	"main-svc/usecases"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m MockUserRepository) FindByID(id int) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}
func (m MockUserRepository) Store(user domain.User) (*domain.User, error) {
	return nil, nil
}

func (m MockUserRepository) FindByUsername(username string) (*domain.User, error) {
	return nil, nil
}

func (m MockUserRepository) Delete(id int) error{}{
	return nil
}

func TestShowUserDataBasedOnID_WithMock(t *testing.T) {
	mock := new(MockUserRepository)

	userInteractor := usecases.UserInteractor{
		UserRepository: mock,
	}

	t.Run("Should return user data when given id", func(t *testing.T) {
		mock.On("FindByID", 1).Return(&domain.User{
			ID:       1,
			Name:     "John Doe",
			Username: "johndoe",
		}, nil)
		userData, _ := userInteractor.ShowUserDataBasedOnID(1)

		want := usecases.User{
			ID:       1,
			Name:     "John Doe",
			Username: "johndoe",
		}

		got := usecases.User{
			ID:       userData.ID,
			Name:     userData.Name,
			Username: userData.Username,
		}

		assert.Equal(t, want, got)
	})

	t.Run("Should return error when given user id is not found", func(t *testing.T) {
		mock.On("FindByID", 69).Return(&domain.User{}, domain.ErrorUserNotFound)
		_, err := userInteractor.ShowUserDataBasedOnID(69)

		assert.ErrorIs(t, err, domain.ErrorUserNotFound)
	})
}
