package usecases

import (
	"main-svc/domain"
)

type UserInteractor struct {
	UserRepository domain.UserRepository
}

type User struct {
	ID       int
	Name     string
	Username string
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
