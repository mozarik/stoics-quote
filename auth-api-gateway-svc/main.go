package main

import (
	"mini-project/auth-api-gateway-svc/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const ROOT_DIR = "../"

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.Logger())

	r := route.NewRoute()
	e.GET("/pingmain", r.PingMainService)
	e.Logger.Fatal(e.Start(":3002"))
}
