package interfaces

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type QuoteGetter interface {
	GetQuoteRadom() (Quote, error)
}

type Quote struct {
	ID          int    `json:"id"`
	Body        string `json:"body"`
	Author      string `json:"author"`
	QuoteSource string `json:"quotesource"`
}

func (q Quote) GetQuoteRadom() (Quote, error) {
	url := "http://stoic-server.herokuapp.com/random"
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Cannot getting response from %v %v", url, err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("cannot read response data to []byte error: %v", err)
	}

	var quote Quote
	err = json.Unmarshal(responseData, &quote)
	if err != nil {
		log.Printf("cannot unmarshal response data to Quote struct error: %v", err)
	}

	return quote, err
}
