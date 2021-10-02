package main

import (
	"api-gateway-svc/route"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// SETUP ALL ENV
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// The port that the service will be listening on
	SERVICE_PORT := os.Getenv("SERVICE_PORT")
	// BaseURL of the MAIN_SERVICE
	MAIN_SERVICE_BASE_URL := os.Getenv("MAIN_SERVICE_BASE_URL")

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.Logger())

	r := route.NewRoute(
		MAIN_SERVICE_BASE_URL,
	)
	e.GET("/pingmain", r.PingMainService)
	e.Logger.Fatal(e.Start(":" + SERVICE_PORT))
}
