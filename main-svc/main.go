package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Data struct {
	Message string
}

func main() {

	// PUT ALL ENV HERE
	SERVICE_PORT := os.Getenv("SERVICE_PORT")

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, Data{
			Message: "pong from main svc",
		})
	})
	e.Logger.Fatal(e.Start(":" + SERVICE_PORT))
}
