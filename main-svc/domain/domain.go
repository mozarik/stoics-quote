package domain

import "errors"

var (
	ErrorUserNotFound = errors.New("user with corresponding id is not found")
)

type UserRepository interface {
	FindByID(id int) (*User, error)
	Store(user User) (*User, error)
	Delete(id int) error
}

type User struct {
	ID       int
	Name     string
	Username string
	Password string
}
