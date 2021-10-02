package route

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Route struct {
}

type Data struct {
	Message string `json:"message"`
}

func NewRoute() *Route {
	return &Route{}
}

func (r Route) PingMainService(c echo.Context) error {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:3001/ping", nil)
	if err != nil {
		return c.JSON(http.StatusNoContent, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusNoContent, err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusNoContent, err)
	}
	var data Data
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return c.JSON(http.StatusNoContent, err)
	}

	return c.JSON(200, data)
}
