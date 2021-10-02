package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Data struct {
	Message string
}

func main() {
	// The port the service will be listening to
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	SERVICE_PORT := os.Getenv("SERVICE_PORT")
	fmt.Println(SERVICE_PORT)
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, Data{
			Message: "pong from main svc",
		})
	})
	e.Logger.Fatal(e.Start(":" + SERVICE_PORT))
}
