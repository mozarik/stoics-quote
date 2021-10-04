package domain

type QuoteRepository interface {
	Save(quote Quote) error
	FindUserFavorites(userID int) ([]Quote, error)
}

type Quote struct {
	ID          int
	Body        string
	Author      string
	QuoteSource string
}
