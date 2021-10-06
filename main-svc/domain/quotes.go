package domain

type QuoteRepository interface {
	FindByID(id int) (*Quote, error)
	Save(quote Quote) error
	FindUserFavorites(userID int) ([]Quote, error)
	// SaveUserFavorite is used to save the data to link table
	SaveUserFavorite(userID int, quoteID int) error
}

type Quote struct {
	ID          int
	Body        string
	Author      string
	QuoteSource string
}
