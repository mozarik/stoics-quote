package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type QuoteGetter interface {
	GetQuoteResponseBody() ([]byte, error)
}

type Quote struct {
	ID          int    `json:"id"`
	Body        string `json:"body"`
	Author      string `json:"author"`
	QuoteSource string `json:"quotesource"`
}

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

type QuoteGetterImpl struct {
	QuoteGetter QuoteGetter
}

func RandomNumberRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	n := min + rand.Intn(max-min+1)
	return n
}

// This function get the Quote from this endpoint
// https://stoic-server.herokuapp.com/quotes/{:id}
// We try to randomize the ID
func GetQuoteFromThirdPartyAPI() (Quote, error) {
	id := RandomNumberRange(1, 1000)
	url := fmt.Sprintf("https://stoic-server.herokuapp.com/quotes/%d", id)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("cannot create request error: %v", err)
		return Quote{}, err
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Printf("cannot do request error: %v", err)
		return Quote{}, err
	}
	defer response.Body.Close()

	var quote []Quote
	err = json.NewDecoder(response.Body).Decode(&quote)
	if err != nil {
		log.Printf("cannot decode response data to Quote struct error: %v", err)
		return Quote{}, err
	}

	return quote[0], nil
}
