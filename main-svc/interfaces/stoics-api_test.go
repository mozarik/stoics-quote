package interfaces_test

import (
	"encoding/json"
	"io/ioutil"
	"main-svc/interfaces"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQuote struct {
	mock.Mock
}

func (m MockQuote) GetQuoteRadom() (interfaces.Quote, error) {
	args := m.Called()
	return args.Get(0).(interfaces.Quote), args.Error(1)
}

func TestGetQuotesByID(t *testing.T) {
	url := "http://stoic-server.herokuapp.com/quotes/1"
	response, err := http.Get(url)
	assert.NoError(t, err)

	responseData, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)

	var quote []interfaces.Quote
	err = json.Unmarshal(responseData, &quote)
	assert.NoError(t, err)

	want := interfaces.Quote{
		ID: 1,
		Body: "There’s no difference between the one and the other - " +
			"you didn’t exist and you won’t exist - you’ve got no concern with either period. " +
			"As it is with a play, so it is with life - what matters is not how long the acting " +
			"lasts, but how good it is. It is not important at what point you stop. Stop wherever you will " +
			"- only make sure that you round it off with a good ending.",
		Author:      "Seneca",
		QuoteSource: "Letter LXXVII",
	}

	got := quote[0]

	assert.EqualValues(t, want, got)
}
