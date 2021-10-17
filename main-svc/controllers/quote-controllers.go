package controllers

import (
	"main-svc/domain"
	"main-svc/interfaces"
	"main-svc/middleware"
	"main-svc/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type QuoteUsecase interface {
	GetQuote() (interfaces.Quote, error)
	UserSaveFavoriteQuote(userID int, quoteData domain.Quote) error
	ListAllFavoriteQuotes(userID int) (usecases.UserFavoriteQuotes, error)
}

type QuoteController struct {
	QUsecase QuoteUsecase
}

func NewQuoteController(QUsecase QuoteUsecase) *QuoteController {
	return &QuoteController{QUsecase: QUsecase}
}

func (qc *QuoteController) Quote(c echo.Context) error {
	quote, err := qc.QUsecase.GetQuote()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, quote)
}

func (qc *QuoteController) SaveFavoriteQuote(c echo.Context) error {
	// /users/:userId
	quoteID, err := strconv.Atoi(c.Param("quoteId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	userID := middleware.GetUser(c).ID
	quoteData, err := qc.QUsecase.GetQuote()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	quoteDataDomain := domain.Quote{
		ID:          quoteID,
		Body:        quoteData.Body,
		Author:      quoteData.Author,
		QuoteSource: quoteData.QuoteSource,
	}
	err = qc.QUsecase.UserSaveFavoriteQuote(userID, quoteDataDomain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Success saved quote as favorite"})
}

func (qc *QuoteController) ListAllFavoriteQuotes(c echo.Context) error {
	userID := middleware.GetUser(c).ID
	userFavoriteQuotes, err := qc.QUsecase.ListAllFavoriteQuotes(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, userFavoriteQuotes)
}
