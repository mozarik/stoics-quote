package usecases

import (
	"log"
	"main-svc/domain"
	"main-svc/interfaces"
)

type UserInteractor struct {
	UserRepository domain.UserRepository
}

type QuoteInteractor struct {
	UserInteractor  UserInteractor
	QuoteRepository domain.QuoteRepository
	QuoteGetter     interfaces.QuoteGetter
}

func NewUserInteractor(ur domain.UserRepository) *UserInteractor {
	return &UserInteractor{
		UserRepository: ur,
	}
}

func NewQuoteInteractor(ur UserInteractor, qr domain.QuoteRepository, qg interfaces.QuoteGetter) *QuoteInteractor {
	return &QuoteInteractor{
		UserInteractor:  ur,
		QuoteRepository: qr,
		QuoteGetter:     qg,
	}
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

func (i UserInteractor) CreateUser(user domain.User) (User, error) {
	var data User
	userData, err := i.UserRepository.Store(user)
	if err != nil {
		return data, domain.ErrorUserAlreadyExists
	}
	data.ID = userData.ID
	data.Name = userData.Name
	data.Username = userData.Username
	return data, nil
}

// Check if userID is in the database
func (i UserInteractor) UserExists(userID int) bool {
	_, err := i.UserRepository.FindByID(userID)
	return err == nil
}

// func (q QuoteInteractor) QuoteExists(quoteID int) bool {
// 	_, err := q.QuoteRepository.FindByID(quoteID)
// 	return err == nil
// }

func (q QuoteInteractor) UserSaveFavoriteQuote(userID int, quoteData domain.Quote) error {
	// Check the quote id in the database
	// Use FindByID to check if the quote exists
	// - if not save the quote in the database to quotes table
	quote, err := q.QuoteRepository.FindByID(quoteData.ID)
	if err != nil && quote.ID == 0 {
		log.Println("Quote not found, save it to the database then")
		err = q.QuoteRepository.Save(quoteData)
		if err != nil {
			return err
		}
	}
	// Check if userID is in the database
	// userExist := q.UserInteractor.UserExists(userID)
	checkUserExist := func(userID int) bool {
		_, err := q.UserInteractor.UserRepository.FindByID(userID)
		return err == nil
	}
	userExist := checkUserExist(userID)
	if !userExist {
		return domain.ErrorUserNotFound
	}
	// save the user id and quote id in the database to userfavoritequotes table
	err = q.QuoteRepository.SaveUserFavorite(userID, quoteData.ID)
	if err != nil {
		return err
	}

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

func (q QuoteInteractor) GetQuote() (interfaces.Quote, error) {
	quote, err := q.QuoteGetter.GetQuoteResponseBody()
	if err != nil {
		return interfaces.Quote{}, err
	}
	return quote, nil
}
