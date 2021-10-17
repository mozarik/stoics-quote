package controllers

type UserCreateResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type QuoteResponse struct {
	ID          int    `json:"id"`
	Body        string `json:"body"`
	Author      string `json:"author"`
	QuoteSource string `json:"quote_source"`
}
