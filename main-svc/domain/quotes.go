package domain

type QuoteRepository interface {
	// GetQuote returns a quote by its id
	FindByID(id int) (*Quote, error)
	// Save is to Save the quote to database
	Save(quote Quote) error
	// FindUserFavorites returns a list of quotes that the user has favorited
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
