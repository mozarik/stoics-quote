package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const ROOT_DIR = "../"

type Data struct {
	Message string
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, Data{
			Message: "pong from main svc",
		})
	})
	e.Logger.Fatal(e.Start(":3001"))
}
