package interfaces

import (
	"main-svc/domain"
)

type QuoteRepo DbRepo

// Implement domain.QuoteRepository interface

func (repo QuoteRepo) Save(quote domain.Quote) error {
	return nil
}

func (repo QuoteRepo) FindUserFavorites(userID int) ([]domain.Quote, error) {
	return []domain.Quote{}, nil
}
