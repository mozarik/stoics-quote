package controllers

import (
	"main-svc/domain"
	"main-svc/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Let the caller define the interface

type UserUsecase interface {
	ShowUserDataBasedOnID(userID int) (usecases.User, error)
	CreateUser(user domain.User) (usecases.User, error)
}

// Get user data by ID
type UserController struct {
	userUc UserUsecase
}

func NewUserController(userUc UserUsecase) *UserController {
	return &UserController{userUc: userUc}
}

func (uc *UserController) Register(c echo.Context) error {
	var req UserCreateRequest
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	user, err := uc.userUc.CreateUser(domain.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}

func (uc *UserController) GetUser(c echo.Context) error {
	// /users/:userId
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	userData, err := uc.userUc.ShowUserDataBasedOnID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, userData)
}
