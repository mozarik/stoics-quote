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

func (m MockUserRepository) Delete(id int) error {
	return nil
}

type MockQuoteRepository struct {
	mock.Mock
}

func (m MockQuoteRepository) Save(quote domain.Quote) error {
	args := m.Called(quote)
	return args.Error(0)
}
func (m MockQuoteRepository) FindUserFavorites(userID int) ([]domain.Quote, error) {
	args := m.Called(userID)
	return args.Get(0).([]domain.Quote), args.Error(1)
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

func TestQuoteInteractor_ListAllFavoriteQuotes(t *testing.T) {
	mockQuote := new(MockQuoteRepository)
	mockUser := new(MockUserRepository)

	quoteInteractor := usecases.QuoteInteractor{
		QuoteRepository: mockQuote,
		UserInteractor: usecases.UserInteractor{
			UserRepository: mockUser,
		},
	}
	t.Run("Should return aggregate of all user saved quote when userID is registered (in database)", func(t *testing.T) {
		mockListFavoriteQuotes := []domain.Quote{
			{
				ID:          1,
				Body:        "Quote 1",
				Author:      "Author 1",
				QuoteSource: "Source 1",
			},
			{
				ID:          2,
				Body:        "Quote 2",
				Author:      "Author 2",
				QuoteSource: "Source 1",
			},
		}
		mockUser.On("FindByID", 1).Return(&domain.User{
			ID:       1,
			Name:     "John Doe",
			Username: "johndoe",
		}, nil)
		mockQuote.On("FindUserFavorites", 1).Return(mockListFavoriteQuotes, nil)

		listUserFavQuote, err := quoteInteractor.ListAllFavoriteQuotes(1)
		assert.NoError(t, err)

		want := usecases.UserFavoriteQuotes{
			User: usecases.User{
				ID:       1,
				Name:     "John Doe",
				Username: "johndoe",
			},
			Quote: mockListFavoriteQuotes,
		}

		got := usecases.UserFavoriteQuotes{
			User: usecases.User{
				ID:       listUserFavQuote.User.ID,
				Name:     listUserFavQuote.User.Name,
				Username: listUserFavQuote.User.Username,
			},
			Quote: listUserFavQuote.Quote,
		}

		assert.Equal(t, want, got)

	})
}
