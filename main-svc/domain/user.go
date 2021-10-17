package domain

import "errors"

var (
	ErrorUserNotFound      = errors.New("user with corresponding id is not found")
	ErrorUserAlreadyExists = errors.New("user with corresponding id already exists")
)

type UserRepository interface {
	// This find the user by id and return the user object (with password data)
	FindByID(id int) (*User, error)
	// This store the user by id and return the user object that we save
	Store(user User) (*User, error)
	// This delete the user from the repository by id
	Delete(id int) error
}

type User struct {
	ID       int
	Name     string
	Username string
	Password string
}
