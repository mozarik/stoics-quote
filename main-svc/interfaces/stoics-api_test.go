package interfaces_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"main-svc/interfaces"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetQuotesByID_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping HIT endpoint directly")
	}
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

// We mock Do method of http.Client
type MockDoType func(req *http.Request) (*http.Response, error)

type MockClient struct {
	MockDo MockDoType
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestGetQuotesByID(t *testing.T) {
	jsonResponse := `[
		{
			"id": 1,
			"body": "There’s no difference between the one and the other - you didn’t exist and you won’t exist - you’ve got no concern with either period. As it is with a play, so it is with life - what matters is not how long the acting lasts, but how good it is. It is not important at what point you stop. Stop wherever you will - only make sure that you round it off with a good ending.",
			"author": "Seneca",
			"quote_source": "Letter LXXVII"
		}
	]`

	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	interfaces.Client = &MockClient{
		MockDo: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	result, err := interfaces.GetQuoteFromThirdPartyAPI()
	assert.NoErrorf(t, err, "Failed getting the result")

	if result.ID == 0 {
		assert.Fail(t, "Result is empty")
	}
}
