package usecases

import (
	"main-svc/domain"
)

type UserInteractor struct {
	UserRepository domain.UserRepository
}

type QuoteInteractor struct {
	UserInteractor  UserInteractor
	QuoteRepository domain.QuoteRepository
}

type User struct {
	ID       int
	Name     string
	Username string
}

func newUserFromDomain(user *domain.User) User {
	return User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
	}
}

type UserFavoriteQuotes struct {
	User  User
	Quote []domain.Quote
}

func (i UserInteractor) ShowUserDataBasedOnID(userID int) (User, error) {
	var userData User
	user, err := i.UserRepository.FindByID(userID)
	if err != nil {
		return userData, domain.ErrorUserNotFound
	}
	userData.ID = user.ID
	userData.Name = user.Name
	userData.Username = user.Username
	return userData, nil
}

// Check if userID is in the database
func (i UserInteractor) UserExists(userID int) bool {
	_, err := i.UserRepository.FindByID(userID)
	return err == nil
}

func (q QuoteInteractor) UserSaveFavoriteQuote(userID int, quoteID int) error {
	// Check the quote id in the database
	// - if not save the quote in the database to quotes table
	// save the user id and quote id in the database to userfavoritequotes table
	return nil
}

func (q QuoteInteractor) ListAllFavoriteQuotes(userID int) (UserFavoriteQuotes, error) {
	// Check if userID is in the database
	userExist := q.UserInteractor.UserExists(userID)
	if !userExist {
		return UserFavoriteQuotes{}, domain.ErrorUserNotFound
	}
	// Get the user data based on id
	ud, err := q.UserInteractor.UserRepository.FindByID(userID)
	if err != nil {
		return UserFavoriteQuotes{}, domain.ErrorUserNotFound
	}
	// Format user from domain to usecase
	userData := newUserFromDomain(ud)
	// Get all quotes from the UserFavoriteQutes table based on userID
	quotes, err := q.QuoteRepository.FindUserFavorites(userID)
	if err != nil {
		return UserFavoriteQuotes{}, err
	}

	// Format/Aggregate data from repository to UserFavoriteQuotes
	aggregate := UserFavoriteQuotes{
		User:  userData,
		Quote: quotes,
	}

	return aggregate, nil
}
