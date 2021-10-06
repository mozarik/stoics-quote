package interfaces_test

import (
	"main-svc/domain"
	"main-svc/interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func deleteAllRecords(t *testing.T, gdb *gorm.DB) {
	t.Helper()
	gdb.Exec("DELETE FROM users")
	gdb.Exec("DELETE FROM quotes")
	gdb.Exec("DELETE FROM userfavoritesquotes")
}

func setup(t *testing.T) {
	t.Helper()
	gdb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	assert.NoError(t, err)

	gdb.Exec("DELETE FROM users")
	gdb.Exec("DELETE FROM quotes")
	gdb.Exec("DELETE FROM userfavoritesquotes")

	// Insert 1 User
	err = gdb.Exec("INSERT INTO users (id, name, username, password) VALUES (@id, @name, @username, @password)", map[string]interface{}{
		"id":       1,
		"name":     "John Doe",
		"username": "johndoe",
		"password": "password",
	}).Error
	assert.NoError(t, err)

	// Insert 1 more user
	err = gdb.Exec("INSERT INTO users (id, name, username, password) VALUES (@id, @name, @username, @password)", map[string]interface{}{
		"id":       2,
		"name":     "John Doe2",
		"username": "johndoe",
		"password": "password",
	}).Error

	assert.NoError(t, err)
	// Insert 2 quotes
	err = gdb.Exec("INSERT INTO quotes (id, body, author, quote_source) VALUES (@id, @body, @author, @quote_source)", map[string]interface{}{
		"id":           1,
		"body":         "I'm a quote",
		"author":       "John Doe",
		"quote_source": "john books",
	}).Error
	assert.NoError(t, err)

	err = gdb.Exec("INSERT INTO quotes (id, body, author, quote_source) VALUES (@id, @body, @author, @quote_source)", map[string]interface{}{
		"id":           2,
		"body":         "I'm another quote",
		"author":       "John Doe",
		"quote_source": "john books",
	}).Error
	assert.NoError(t, err)

	err = gdb.Exec("INSERT INTO userfavoritesquotes (user_id, quote_id) VALUES (@user_id, @quote_id)", map[string]interface{}{
		"user_id":  1,
		"quote_id": 1,
	}).Error
	assert.NoError(t, err)

	err = gdb.Exec("INSERT INTO userfavoritesquotes (user_id, quote_id) VALUES (@user_id, @quote_id)", map[string]interface{}{
		"user_id":  1,
		"quote_id": 2,
	}).Error
	assert.NoError(t, err)

	t.Cleanup(func() {
		gdb.Exec("DELETE FROM users")
		gdb.Exec("DELETE FROM quotes")
		gdb.Exec("DELETE FROM userfavoritesquotes")
	})
}

func TestFindByID(t *testing.T) {
	gdb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	assert.NoError(t, err)

	quoteRepo := interfaces.QuoteRepo{
		DB: gdb,
	}

	setup(t)

	t.Run("SUCCESSFULL PATH: return quote by the quote id in quotes table", func(t *testing.T) {
		quote, err := quoteRepo.FindByID(1)
		assert.NoError(t, err)

		want := &domain.Quote{
			ID:          1,
			Body:        "I'm a quote",
			Author:      "John Doe",
			QuoteSource: "john books",
		}

		got := quote

		assert.Equal(t, want, got)

	})

	t.Run("FAILURE PATH: return error if quote id is not found in quotes table", func(t *testing.T) {
		_, err := quoteRepo.FindByID(3)
		assert.Error(t, err)
	})
}

func TestFindUserFavoriteQuote(t *testing.T) {
	gdb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	assert.NoError(t, err)

	quoteRepo := interfaces.QuoteRepo{
		DB: gdb,
	}

	setup(t)

	quotes, err := quoteRepo.FindUserFavorites(1)
	assert.NoError(t, err)

	want := []domain.Quote{
		{
			ID:          1,
			Body:        "I'm a quote",
			Author:      "John Doe",
			QuoteSource: "john books",
		},
		{
			ID:          2,
			Body:        "I'm another quote",
			Author:      "John Doe",
			QuoteSource: "john books",
		},
	}

	assert.EqualValues(t, want, quotes)

}

func TestQuoteRepo_Save(t *testing.T) {
	gdb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	assert.NoError(t, err)

	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		quote domain.Quote
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Successful save",
			fields: fields{
				DB: gdb,
			},
			args: args{
				quote: domain.Quote{
					ID:          1,
					Body:        "I'm a quote",
					Author:      "John Doe",
					QuoteSource: "john books",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := interfaces.QuoteRepo{
				DB: tt.fields.DB,
			}
			if err := repo.Save(tt.args.quote); (err != nil) != tt.wantErr {
				t.Errorf("QuoteRepo.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			deleteAllRecords(t, tt.fields.DB)
		})
	}
}

func TestSaveUserFavorite(t *testing.T) {
	gdb, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	assert.NoError(t, err)

	quoteRepo := interfaces.QuoteRepo{
		DB: gdb,
	}

	setup(t)

	t.Run("Successful save user favorite to link table (userfavoritesquotes)", func(t *testing.T) {
		err := quoteRepo.SaveUserFavorite(2, 1)
		assert.NoError(t, err)

		// Make sure the record was inserted
		var tempStruct struct {
			UserID  int
			QuoteID int
		}
		row := gdb.Raw("SELECT user_id, quote_id FROM userfavoritesquotes WHERE user_id = ? AND quote_id = ?", 2, 1).Row()
		err = row.Scan(&tempStruct.UserID, &tempStruct.QuoteID)
		assert.NoError(t, err)

		assert.Equal(t, 2, tempStruct.UserID)
		assert.Equal(t, 1, tempStruct.QuoteID)

	})

}
