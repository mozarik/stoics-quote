package domain

import "errors"

var (
	ErrorUserNotFound = errors.New("user with corresponding id is not found")
)

type UserRepository interface {
	FindByID(id int) (*User, error)
	Store(user User) error
}

type User struct {
	ID       int
	Name     string
	Username string
	Password string
}
