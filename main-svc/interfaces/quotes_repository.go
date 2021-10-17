package interfaces

import (
	"database/sql"
	"main-svc/domain"

	"gorm.io/gorm"
)

type QuoteRepo DbRepo

func NewQuoteRepo(db *gorm.DB) QuoteRepo {
	return QuoteRepo{DB: db}
}

// Implement domain.QuoteRepository interface
func (repo QuoteRepo) Save(quote domain.Quote) error {
	query := `INSERT INTO quotes (id, body, author, quote_source) VALUES (@id, @body, @author, @quote_source)`
	err := repo.DB.Exec(query,
		sql.Named("id", quote.ID),
		sql.Named("body", quote.Body),
		sql.Named("author", quote.Author),
		sql.Named("quote_source", quote.QuoteSource)).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo QuoteRepo) FindByID(id int) (*domain.Quote, error) {
	query := `SELECT id, body, author, quote_source FROM quotes WHERE id = @id`

	var quote domain.Quote
	row := repo.DB.Debug().Raw(query, sql.Named("id", id)).Row()
	err := row.Scan(&quote.ID, &quote.Body, &quote.Author, &quote.QuoteSource)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}

func (repo QuoteRepo) SaveUserFavorite(userID int, quoteID int) error {
	query := `INSERT INTO userfavoritesquotes (user_id, quote_id) VALUES (@userID, @quoteID)`
	err := repo.DB.Exec(query, sql.Named("userID", userID), sql.Named("quoteID", quoteID)).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo QuoteRepo) FindUserFavorites(userID int) ([]domain.Quote, error) {
	query := `SELECT q.id, q.body, q.author, q.quote_source FROM quotes q
	JOIN userfavoritesquotes on (q.id = userfavoritesquotes.quote_id)
	JOIN users ON (users.id = userfavoritesquotes.user_id)
	WHERE users.id = @userID`

	rows, err := repo.DB.Debug().Raw(query, sql.Named("userID", userID)).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quotes []domain.Quote

	for rows.Next() {
		var quote domain.Quote
		if err := rows.Scan(&quote.ID, &quote.Body, &quote.Author, &quote.QuoteSource); err != nil {
			return quotes, err
		}
		quotes = append(quotes, quote)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return quotes, nil
}
