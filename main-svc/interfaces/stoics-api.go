package interfaces

import (
	"encoding/json"
	"fmt"
	"log"
	"main-svc/domain"
	"math/rand"
	"net/http"
	"time"
)

type QuoteGetter interface {
	GetQuoteResponseBody() (Quote, error)
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
	QuoteSaver QuoteSaver
}

type QuoteSaver interface {
	SaveQuote(quote domain.Quote) error
}

type QuoteSaverImpl struct {
	QRepo domain.QuoteRepository
}

func NewQuoteGetter(qs QuoteSaver) QuoteGetter {
	return &QuoteGetterImpl{
		QuoteSaver: qs,
	}
}

func ProvideQuoteSaver(qr domain.QuoteRepository) QuoteSaver {
	return &QuoteSaverImpl{
		QRepo: qr,
	}
}

func (qs *QuoteSaverImpl) SaveQuote(quoteData domain.Quote) error {
	quote, err := qs.QRepo.FindByID(quoteData.ID)
	if err != nil && quote.ID == 0 {
		log.Println("Quote not found, save it to the database then")
		err = qs.QRepo.Save(quoteData)
		if err != nil {
			return err
		}
	}
	return nil
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

func (qg *QuoteGetterImpl) GetQuoteResponseBody() (Quote, error) {
	quote, err := GetQuoteFromThirdPartyAPI()
	if err != nil {
		log.Printf("cannot get quote from third party api error: %v", err)
		return quote, err
	}
	quoteToDomain := domain.Quote{
		ID:          quote.ID,
		Body:        quote.Body,
		Author:      quote.Author,
		QuoteSource: quote.QuoteSource,
	}
	err = qg.QuoteSaver.SaveQuote(quoteToDomain)
	if err != nil {
		log.Printf("cannot save quote to database error: %v", err)
		return quote, err
	}
	return quote, nil
}
